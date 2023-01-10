# nfv-lab-cds
Repositorio Helm para la Práctica 4 de RDSV con SDA+CDS

![nfv-lab-cds](https://github.com/martinezgarciadavid/nfv-lab-cds/blob/main/nfv-lab-cds.drawio.png)

## Quick deployment and testing guide

1. Launch RDSV-K8S and RDSV-OSM virtual machines.

NOTE: Bear in mind that these virtual machines may need more resources than the ones originally defined. A basic recommendation would be to set these machines to use, at least, 8 CPU cores and 12 GB of RAM.

2. On both machines open a terminal, clone the repository and change to its directory.

```
$ cd Desktop/
$ git clone https://github.com/martinezgarciadavid/nfv-lab-cds.git
$ cd nfv-lab-cds/
```

3. On RDSV-K8S, start the VNX scenarios.

```
$ cd vnx/
$ ./vnx_create_all.sh
```

4. On RDSV-OSM, from the GUI:
- Open Mozilla Firefox and access the following URL: http://localhost. Login to the OSM GUI using the following credentials: admin/admin.
