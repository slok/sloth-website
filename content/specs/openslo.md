---
title: "OpenSLO"
weight: 030
---

{{< hint warning >}}
**Only CLI, Not Kubernetes compatible**\
This spec can't be used with Kubernetes CRDs, use [Kubernetes spec]({{< relref "./kubernetes.md" >}}) instead.
{{< /hint >}}

Sloth supports [OpenSLO v1alpha](https://github.com/OpenSLO/OpenSLO) spec, however we need to take into account that it has some restrictions:

- OpenSLO time window restricted to 30 days.
- Only Objective ratio metrics are supported.
- Only Prometheus and PromQL query types are supported.
- Configuration fields not required by Sloth will be ignored.

Regarding Sloth features, [OpenSLO] spec doesn't support all of the sloth features:

- No Prometheus labels support.
- No alerting support.
- No SLI plugins support.
- No Kubernetes support (at least until official OpenSLO CRDs are released).

Check [Examples]({{< relref "../examples/openslo" >}})

CLI example:

```bash
sloth generate -i ./examples/openslo-getting-started.yml  -o /tmp/openslo-getting-started.yml
INFO[0000] SLI plugins loaded                            plugins=0 svc=storage.FileSLIPlugin version=v0.6.0-8-ga8f37a2
INFO[0000] Generating from OpenSLO spec                  version=v0.6.0-8-ga8f37a2
INFO[0000] Multiwindow-multiburn alerts generated        out=/tmp/openslo-getting-started.yml slo=my-service-sloth-slo-my-service-0 svc=generate.prometheus.Service version=v0.6.0-8-ga8f37a2
INFO[0000] SLI recording rules generated                 out=/tmp/openslo-getting-started.yml rules=8 slo=my-service-sloth-slo-my-service-0 svc=generate.prometheus.Service version=v0.6.0-8-ga8f37a2
INFO[0000] Metadata recording rules generated            out=/tmp/openslo-getting-started.yml rules=7 slo=my-service-sloth-slo-my-service-0 svc=generate.prometheus.Service version=v0.6.0-8-ga8f37a2
INFO[0000] SLO alert rules generated                     out=/tmp/openslo-getting-started.yml rules=0 slo=my-service-sloth-slo-my-service-0 svc=generate.prometheus.Service version=v0.6.0-8-ga8f37a2
INFO[0000] Prometheus rules written                      format=yaml groups=2 out=/tmp/openslo-getting-started.yml svc=storage.IOWriter version=v0.6.0-8-ga8f37a2
```

[openslo]: https://openslo.com/
