FROM ubuntu:20.04
# install required packages
RUN apt-get clean
RUN apt-get update \
    && apt-get dist-upgrade -y \
    && apt-get install -y \
    net-tools \
    traceroute \
    curl \
    iptables \
    inetutils-ping \
    nano \
    build-essential \
    bridge-utils \
    isc-dhcp-server \
    tcpdump \
    openvswitch-switch \
    openvswitch-common \
    iperf3 \
    iproute2 \
    vim \
    softflowd \
    && apt-get autoremove -y \
    && apt-get autoclean

COPY vnx_config_nat /usr/bin/
COPY isc-dhcp-server /etc/default/isc-dhcp-server
COPY dhcpd.conf /etc/dhcp/dhcpd.conf 
COPY softflowd /etc/default/softflowd
COPY softflowd /etc/softflowd/default.conf
COPY capture+crypto+type2+8+vlan.pcap /etc/softflowd/capture+crypto+type2+8+vlan.pcap
