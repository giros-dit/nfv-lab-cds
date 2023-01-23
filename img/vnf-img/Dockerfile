FROM ubuntu:20.04

# Variables to automatically configure tzdata
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Madrid

# Install required packages
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
    git \
    dh-autoreconf \
    gcc \
    make \
    perl \
    libpcap-dev \
    && apt-get autoremove -y \
    && apt-get autoclean

# Download, compile and install softflowd from sources (GitHub repository)    
RUN git clone https://github.com/irino/softflowd.git && cd softflowd && autoreconf -if && ./configure && make && make install

# Directory to store PCAP files to inject
RUN mkdir -p /usr/local/etc/softflowd/

# Copy configuration files and PCAP file with cryptomining traffic
COPY vnx_config_nat /usr/bin/
COPY isc-dhcp-server /etc/default/isc-dhcp-server
COPY dhcpd.conf /etc/dhcp/dhcpd.conf 
# CURRENTLY NOT USED -> COPY softflowd /etc/default/softflowd
# CURRENTLY NOT USED -> COPY softflowd /etc/softflowd/default.conf
COPY capture+crypto+type2+8+vlan.pcap /usr/local/etc/softflowd/capture+crypto+type2+8+vlan.pcap

