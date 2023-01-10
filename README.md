# nfv-lab-cds
Helm repository based on a [NFV laboratory practice](https://github.com/educaredes/nfv-lab) with [SDA+CDS](https://github.com/giros-dit/semantic-data-aggregator) integration, for demostrative purposes.

![nfv-lab-cds](https://github.com/martinezgarciadavid/nfv-lab-cds/blob/main/nfv-lab-cds.drawio.png)

## Quick deployment and testing guide

1. Launch _RDSV-K8S_ and _RDSV-OSM_ virtual machines.

**NOTE**: Bear in mind that these virtual machines need more resources than the ones originally defined. A basic recommendation would be to set these machines to use, at least, 8 CPU cores and 12 GB of RAM.
For the _RDSV-K8S_ machine, install any required software, like _Helm_ or the standalone _kubectl_ util (and point it to the _MicroK8s_ cluster). Also, your Docker daemon must be configured to use a private repository, and this repository must contain all the custom images of the services that will be deployed. The following links give the instructions to accomplish this:
- [Running Flink clusters in application mode on Kubernetes](https://github.com/giros-dit/semantic-data-aggregator/tree/develop/kubernetes/flink-operator).
- [Dockerfiles of services](https://github.com/giros-dit/semantic-data-aggregator/tree/develop/kubernetes/flink-operator/flink-cluster-templates/netflow).
- [Working with a private Docker registry](https://microk8s.io/docs/registry-private).
- [Configuring the standalone _kubectl_ tool to point to a _MicroK8s_ cluster](https://microk8s.io/docs/working-with-kubectl).

2. On both machines open a terminal, clone the repository and change to its directory.

```
$ cd Desktop/
$ git clone https://github.com/martinezgarciadavid/nfv-lab-cds.git
$ cd nfv-lab-cds/
```

3. On _RDSV-K8S_, start the VNX scenarios.

```
$ cd vnx/
$ ./vnx_create_all.sh
```

4. On _RDSV-K8S_, launch all SDA-related services.
```
$ cd ../kubernetes/
$ cd services/
$ kubectl apply -f zookeeper.yaml
$ kubectl apply -f kafka.yaml
$ cd ../netflow-full-pipeline/
$ helm install -f values.yaml netflow-full-pipeline .
```

5. On _RDSV-OSM_, add this _Helm_ chart repository to OSM. To do so, open _Mozilla Firefox_ on this machine, open http://localhost and login to the OSM GUI with the following credentials (admin/admin). Once logged in, navigate to the menu on the left-side of the interface and go to _K8s_ -> _K8s Repos_. Then, click on the button _Add K8s repository_ and fill in the form with the following information:
```
Name: helmchartrepo
Type: Helm Chart
URL: https://martinezgarciadavid.github.io/nfv-lab-cds
Description: Type a short, descriptive phrase.
```
