---
title: "Default (prometheus/v1)"
weight: 010
---

{{< hint warning >}}
**Only CLI, Not Kubernetes compatible**\
This spec can't be used with Kubernetes CRDs, use [Kubernetes spec]({{< relref "./kubernetes.md" >}}) instead.
{{< /hint >}}

This is the default format for Sloth, doesn't depend on Kubernetes CRDs or anything.

Normally this format is used when the integration you want is a vanilla Prometheus integration with raw Prometheus rules, so, this will generate the prometheus [recording][prom-recordings] and [alerting][prom-alerts] rules in Standard Prometheus YAML format.

- [Spec]
- [Examples]({{< relref "../examples/default" >}})

CLI Example:

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

[spec]: https://pkg.go.dev/github.com/slok/sloth@v0.6.0/pkg/prometheus/api/v1
[prom-recordings]: https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/
[prom-alerts]: https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/
