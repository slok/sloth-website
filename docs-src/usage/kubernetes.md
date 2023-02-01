---
title: "Kubernetes controller"
weight: 020
---

`kubernetes-controller` command runs Sloth as a controller/operator that will react on [`sloth.slok.dev/v1/PrometheusServiceLevel`][sloth-crd] CRD.

In the controller mode, Sloth will be running listneing to Kubernetes Sloth SLO CR events and generate the required prometheus-operator [CR rules][prom-op-rules]

In the end, the controller mode makes the same work as the CLI however integrates better with a native Kubernetes flow.

!!! info "Sloth [CRD][sloth-crd] is required"

    ```bash
    $ kubectl apply -f ./pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
    ```

!!! info "Prometheus Operator Rules [CRD][prom-op-rules] is required"

    ```bash
    $ kubectl apply -f ./test/integration/crd/prometheus-operator-crd.yaml
    ```

!!! tip "Disable optimized rules"

    By default Sloth will try optimizing long time windows (e.g 30 day) by sacrificing data precision. If you want to disable this use `--disable-optimized-rules`.  

## Deploying in Kubernetes

Check [Kubernetes installing section](../introduction/install.md#kubernetes)

!!! example

    ```bash
    # Run sloth controller.
    $ kubectl create ns monitoring
    $ kubectl apply -f ./deploy/kubernetes/raw/sloth.yaml

    # Deploy some SLOs.
    $ kubectl apply -f ./examples/k8s-getting-started.yml

    # Get CRs.
    $ kubectl -n monitoring get slos
    NAME                   SERVICE           DESIRED SLOS   READY SLOS   GEN OK   GEN AGE   AGE
    sloth-slo-my-se rvice   myservice         1              1            true     27s       27s

    $ kubectl -n monitoring get prometheusrules
    NAME                  AGE
    sloth-slo-my-service  38s
    ```

[sloth-crd]: https://raw.githubusercontent.com/slok/sloth/main/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
[prom-op-rules]: https://github.com/prometheus-operator/kube-prometheus/blob/main/manifests/setup/0prometheusruleCustomResourceDefinition.yaml
