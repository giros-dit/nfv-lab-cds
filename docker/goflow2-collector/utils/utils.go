package utils

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	decoder "goflow2-collector/decoders"
	"goflow2-collector/decoders/netflow"
	flowmessage "goflow2-collector/pb"
	"goflow2-collector/producer"

	reuseport "github.com/libp2p/go-reuseport"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/yaml.v2"
)

type ProducerConfig *producer.ProducerConfig

func LoadMapping(f io.Reader) (ProducerConfig, error) {
	config := &producer.ProducerConfig{}
	dec := yaml.NewDecoder(f)
	err := dec.Decode(config)
	return config, err
}

func GetServiceAddresses(srv string) (addrs []string, err error) {
	_, srvs, err := net.LookupSRV("", "", srv)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Service discovery: %v\n", err))
	}
	for _, srv := range srvs {
		addrs = append(addrs, net.JoinHostPort(srv.Target, strconv.Itoa(int(srv.Port))))
	}
	return addrs, nil
}

type Logger interface {
	Printf(string, ...interface{})
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
	Warn(...interface{})
	Error(...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Fatalf(string, ...interface{})
}

type BaseMessage struct {
	Src     net.IP
	Port    int
	Payload []byte

	SetTime  bool
	RecvTime time.Time
}

type Transport interface {
	Send([]*flowmessage.FlowMessage)
}

type Formatter interface {
	Format([]*flowmessage.FlowMessage)
}

/*
type DefaultLogTransport struct {
}

func (s *DefaultLogTransport) Publish(msgs []*flowmessage.FlowMessage) {
	for _, msg := range msgs {
		fmt.Printf("%v\n", FlowMessageToString(msg))
	}
}

type DefaultJSONTransport struct {
}

func (s *DefaultJSONTransport) Publish(msgs []*flowmessage.FlowMessage) {
	for _, msg := range msgs {
		fmt.Printf("%v\n", FlowMessageToJSON(msg))
	}
}
*/
type DefaultErrorCallback struct {
	Logger Logger
}

func (cb *DefaultErrorCallback) Callback(name string, id int, start, end time.Time, err error) {
	if _, ok := err.(*netflow.ErrorTemplateNotFound); ok {
		return
	}
	if cb.Logger != nil {
		cb.Logger.Errorf("Error from: %v (%v) duration: %v. %v", name, id, end.Sub(start), err)
	}
}

func UDPRoutine(name string, decodeFunc decoder.DecoderFunc, workers int, addr string, port int, sockReuse bool, logger Logger) error {
	return UDPStoppableRoutine(make(chan struct{}), name, decodeFunc, workers, addr, port, sockReuse, logger)
}

// UDPStoppableRoutine runs a UDPRoutine that can be stopped by closing the stopCh passed as argument
func UDPStoppableRoutine(stopCh <-chan struct{}, name string, decodeFunc decoder.DecoderFunc, workers int, addr string, port int, sockReuse bool, logger Logger) error {
	ecb := DefaultErrorCallback{
		Logger: logger,
	}

	decoderParams := decoder.DecoderParams{
		DecoderFunc:   decodeFunc,
		DoneCallback:  DefaultAccountCallback,
		ErrorCallback: ecb.Callback,
	}

	processor := decoder.CreateProcessor(workers, decoderParams, name)
	processor.Start()

	addrUDP := net.UDPAddr{
		IP:   net.ParseIP(addr),
		Port: port,
	}

	var udpconn *net.UDPConn
	var err error

	if sockReuse {
		pconn, err := reuseport.ListenPacket("udp", addrUDP.String())
		defer pconn.Close()
		if err != nil {
			return err
		}
		var ok bool
		udpconn, ok = pconn.(*net.UDPConn)
		if !ok {
			return err
		}
	} else {
		udpconn, err = net.ListenUDP("udp", &addrUDP)
		if err != nil {
			return err
		}
		defer udpconn.Close()
	}

	payload := make([]byte, 9000)

	localIP := addrUDP.IP.String()
	if addrUDP.IP == nil {
		localIP = ""
	}

	type udpData struct {
		size    int
		pktAddr *net.UDPAddr
		payload []byte
	}

	stopped := atomic.Value{}
	stopped.Store(false)
	udpDataCh := make(chan udpData)
	go func() {
		for {
			u := udpData{}
			u.size, u.pktAddr, _ = udpconn.ReadFromUDP(payload)
			if stopped.Load() == false {
				u.payload = make([]byte, u.size)
				copy(u.payload, payload[0:u.size])
				udpDataCh <- u
			} else {
				return
			}
		}
	}()
	for {
		select {
		case u := <-udpDataCh:
			process(u.size, u.payload, u.pktAddr, processor, localIP, addrUDP, name)
		case <-stopCh:
			stopped.Store(true)
			udpconn.Close()
			close(udpDataCh)
			return nil
		}
	}
}

func process(size int, payload []byte, pktAddr *net.UDPAddr, processor decoder.Processor, localIP string, addrUDP net.UDPAddr, name string) {
	baseMessage := BaseMessage{
		Src:     pktAddr.IP,
		Port:    pktAddr.Port,
		Payload: payload,
	}
	processor.ProcessMessage(baseMessage)

	MetricTrafficBytes.With(
		prometheus.Labels{
			"remote_ip":  pktAddr.IP.String(),
			"local_ip":   localIP,
			"local_port": strconv.Itoa(addrUDP.Port),
			"type":       name,
		}).
		Add(float64(size))
	MetricTrafficPackets.With(
		prometheus.Labels{
			"remote_ip":  pktAddr.IP.String(),
			"local_ip":   localIP,
			"local_port": strconv.Itoa(addrUDP.Port),
			"type":       name,
		}).
		Inc()
	MetricPacketSizeSum.With(
		prometheus.Labels{
			"remote_ip":  pktAddr.IP.String(),
			"local_ip":   localIP,
			"local_port": strconv.Itoa(addrUDP.Port),
			"type":       name,
		}).
		Observe(float64(size))
}
