#!/bin/sh

# Variables necesarias.
KUBECTL="kubectl"
OSMNS="7b2950d8-f92b-4041-9a55-8d1837ad7b0a"
SINAME="renes1"

# Función para obtener el deployment_id de una VNF.
deployment_id() {
    echo `osm ns-list | grep $1 | awk '{split($0,a,"|");print a[3]}' | xargs osm vnf-list --ns | grep $2 | awk '{split($0,a,"|");print a[2]}' | xargs osm vnf-show --literal | grep name | grep $2 | awk '{split($0,a,":");print a[2]}' | sed 's/ //g'`
}

echo "## 1. Obtener deployment_id de VNF:cpe"
OSMCPE=$(deployment_id $SINAME "cpe")

# Referencia a VNF:cpe.
VCPE="deploy/$OSMCPE"

# Base para ejecutar comandos en VNF:cpe.
CPE_EXEC="$KUBECTL exec -n $OSMNS $VCPE --"

echo "## 2. Ejecutando instancia de softflowd e inyectando PCAP con tráfico de criptominado conocido"
$CPE_EXEC /usr/sbin/softflowd -d -r capture+crypto+type2+8+vlan.pcap -v 9 -t maxlife=30s -T ether -n goflow2-1.default.svc.cluster.local:9995 &
sleep 10
