---
title: "CLI generate"
weight: 010
---

`generate` will generate Prometheus rules in different formats based on the specs.

This mode only needs the CLI so its very useful for GitOps, CI, scripts or as a CLI on your toolbox.

`generate` command supports multiple spec types and will detect the input spec type based on the spec, accordingly it will generate the output in different formats depending on the loaded spec:

- [Default]({{< relref "../specs/default.md" >}}) spec generates vanilla Prometheus rules.
- [Kubernetes]({{< relref "../specs/kubernetes.md" >}}) spec generates [Prometheus operator][prometheus-operator] based [CRD Prometheus rules][prom-op-rules].

{{< hint info >}}
**Kubernetes without prometheus-operator**\
In this case, you will need default spec to get vanilla prometheus rules and deploy as you deploy other Prometheus rules in the Kubernetes cluster (e.g. Using configmaps).
{{< /hint >}}

{{< hint info >}}
**CLI mode don't need CRDs**\
Sloth CRD is NOT required in the cluster because the generation happens offline as a CLI. For controller/operator K8s flow, check [Kubernetes controller]({{< relref "./kubernetes.md" >}}) section
{{< /hint >}}

{{< hint warning >}}
**Kubernetes specs need prometheus-operator CRDs**\
Kubernetes specs generate Prometheus operator Rules CRs, this means that the [CRD][prom-op-rules-crd] is required to be registered in the cluster.
{{< /hint >}}

### Example

```bash
$ sloth generate -i ./examples/home-wifi.yml -o /tmp/home-wifi.yml
INFO[0000] Generating from Prometheus spec               version=v0.1.0-43-g5715af5
INFO[0000] Multiwindow-multiburn alerts generated        slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] SLI recording rules generated                 rules=8 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] Metadata recording rules generated            rules=7 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] SLO alert rules generated                     rules=2 slo=home-wifi-good-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] Multiwindow-multiburn alerts generated        slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] SLI recording rules generated                 rules=8 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] Metadata recording rules generated            rules=7 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] SLO alert rules generated                     rules=2 slo=home-wifi-risk-wifi-client-satisfaction svc=generate.prometheus.Service version=v0.1.0-43-g5715af5
INFO[0000] Prometheus rules written                      format=yaml groups=6 out=/tmp/home-wifi.yml svc=storage.IOWriter version=v0.1.0-43-g5715af5
```

[prometheus-operator]: https://github.com/prometheus-operator
[prom-op-rules]: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#prometheusrule
[prom-op-rules-crd]: https://github.com/prometheus-operator/kube-prometheus/blob/main/manifests/setup/prometheus-operator-0prometheusruleCustomResourceDefinition.yaml
