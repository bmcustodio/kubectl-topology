# kubectl-topology

Provides insight into the topology of a Kubernetes cluster.

![Release](https://img.shields.io/github/v/release/bmcstdio/kubectl-topology)
[![Build](https://img.shields.io/travis/com/bmcstdio/kubectl-topology)](https://travis-ci.com/bmcstdio/kubectl-topology)
![License](https://img.shields.io/github/license/bmcstdio/kubectl-topology)

## Introduction

`kubectl-topology` provides insight into the topology of a Kubernetes cluster.
At the moment, this essentially means the distribution of nodes and pods across regions and zones:

```shell
$ kubectl-topology pod --all-namespaces
NAMESPACE     NAME                                     NODE                          REGION         ZONE
kube-system   kube-dns-5f886bf8d8-s7pcc                gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-dns-5f886bf8d8-zhm42                gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-dns-autoscaler-8687c64fc-d5jvm      gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-proxy-gke-gke-1-p-1-50120dfc-26gm   gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-proxy-gke-gke-1-p-1-7023cbca-cz4b   gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
kube-system   kube-proxy-gke-gke-1-p-1-8e1077f6-17st   gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   l7-default-backend-8f479dd9-dkspg        gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   metrics-server-v0.3.1-5c6fbf777-z7jxq    gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
ns-1          nginx                                    gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
```

`kubectl-topology` requires nodes to be adequately labeled with the `topology.kubernetes.io/{region,zone}` labels, or with their deprecated (as of Kubernetes 1.17) equivalents, `failure-domain.beta.kubernetes.io/{region,zone}`. 


## Installation

At the moment, `kubectl-topology` must be installed by running

```shell
$ go get github.com/bmcstdio/kubectl-topology/cmd/kubectl-topology
```

or by cloning this repository, running

```shell
$ make build
```

and copying `./bin/kubectl-topology` to a directory in your `$PATH`.

## Usage & examples

### Node topology

To list nodes by zone and region, run

```shell
$ kubectl-topology node
NAME                          REGION         ZONE
gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
```

To list nodes in a specific zone (e.g. `europe-west1-a`), run

```shell
$ kubectl topology node --zone europe-west1-b
NAME                          REGION         ZONE
gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
```

To list nodes in a specific region (e.g. `europe-west1`), run

```shell
$ kubectl-topology node --region europe-west1
NAME                          REGION         ZONE
gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
```

**NOTE:** `--region` and `--zone` are mutually exclusive.

### Pod topology

To list pods across all namespaces and the region and zones they are in, run

```shell
$ kubectl-topology pod --all-namespaces
NAMESPACE     NAME                                     NODE                          REGION         ZONE
kube-system   kube-dns-5f886bf8d8-s7pcc                gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-dns-5f886bf8d8-zhm42                gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-dns-autoscaler-8687c64fc-d5jvm      gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-proxy-gke-gke-1-p-1-50120dfc-26gm   gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-proxy-gke-gke-1-p-1-7023cbca-cz4b   gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
kube-system   kube-proxy-gke-gke-1-p-1-8e1077f6-17st   gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   l7-default-backend-8f479dd9-dkspg        gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   metrics-server-v0.3.1-5c6fbf777-z7jxq    gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
ns-1          nginx                                    gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
```

To list pods in the `europe-west1-b` zone across all namespaces, run

```shell
$ kubectl-topology pod --all-namespaces --zone europe-west1-b
NAMESPACE     NAME                                     NODE                          REGION         ZONE
kube-system   kube-dns-5f886bf8d8-zhm42                gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-dns-autoscaler-8687c64fc-d5jvm      gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-proxy-gke-gke-1-p-1-8e1077f6-17st   gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   metrics-server-v0.3.1-5c6fbf777-z7jxq    gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
```

To list pods in the `europe-west1` region across all namespaces, run

```shell
$ kubectl-topology pod --all-namespaces --region europe-west1
NAMESPACE     NAME                                     NODE                          REGION         ZONE
kube-system   kube-dns-5f886bf8d8-s7pcc                gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-dns-5f886bf8d8-zhm42                gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-dns-autoscaler-8687c64fc-d5jvm      gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   kube-proxy-gke-gke-1-p-1-50120dfc-26gm   gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   kube-proxy-gke-gke-1-p-1-7023cbca-cz4b   gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
kube-system   kube-proxy-gke-gke-1-p-1-8e1077f6-17st   gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
kube-system   l7-default-backend-8f479dd9-dkspg        gke-gke-1-p-1-50120dfc-26gm   europe-west1   europe-west1-d
kube-system   metrics-server-v0.3.1-5c6fbf777-z7jxq    gke-gke-1-p-1-8e1077f6-17st   europe-west1   europe-west1-b
ns-1          nginx                                    gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
```

To list pods in the `europe-west1-c` zone and in `ns-1` only, run

```shell
$ kubectl-topology pod --namespace ns-1 --zone europe-west1
NAMESPACE   NAME    NODE                          REGION         ZONE
ns-1        nginx   gke-gke-1-p-1-7023cbca-cz4b   europe-west1   europe-west1-c
```

**NOTE:** `--region` and `--zone` are mutually exclusive.

## License

Copyright 2020 bmcstdio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
