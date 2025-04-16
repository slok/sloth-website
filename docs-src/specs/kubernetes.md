---
title: "Kubernetes (CRD)"
weight: 020
---

!!! info ":material-console: Compatible with CLI"
    This spec can be used in the same way the [default spec](./default.md) is used with the CLI.

!!! info ":material-kubernetes: Compatible with Kubernetes"
    This spec can be used as a Kubernetes resource using CR (needs CRD registered on Kubernetes APIserver).

Kubernetes format means that the SLOs use a Kubernetes style spec. Is very similar to the default one, except that this format can be used in two ways, the regular CLI mode, and in Kubernetes controller mode using CRs.

Check related resources:

- [Spec][spec].
- [Examples](../examples/kubernetes/getting-started.md)
- [K8s client][k8s-cli].
- [CRD manifest][crd-manifest].

Will generate from a [Sloth CRD](https://github.com/slok/sloth/tree/main/pkg/kubernetes/api/sloth/v1) spec into [Prometheus-operator][crd rules][prom-op-rules]. This generates the prometheus operator CRDs based on the Sloth CRD template.

!!! example

    ```bash
    $ sloth generate -i ./examples/k8s-home-wifi.yml -o /tmp/k8s-home-wifi.yml
    INFO[0000] Generating from Kubernetes Prometheus spec    version=v0.1.0-43-g5715af5
    INFO[0000] Multiwindow-multiburn alerts generated        slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] SLI recording rules generated                 rules=8 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] Metadata recording rules generated            rules=7 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] SLO alert rules generated                     rules=2 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] Multiwindow-multiburn alerts generated        slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] SLI recording rules generated                 rules=8 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] Metadata recording rules generated            rules=7 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    INFO[0000] SLO alert rules generated                     rules=2 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
    ```

[spec]: https://pkg.go.dev/github.com/slok/sloth/pkg/kubernetes/api/sloth/v1
[k8s-cli]: https://pkg.go.dev/github.com/slok/sloth/pkg/kubernetes/gen/clientset/versioned/typed/sloth/v1
[crd-manifest]: https://raw.githubusercontent.com/slok/sloth/main/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
