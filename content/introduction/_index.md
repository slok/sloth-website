---
title: "Getting started"
weight: 001
---

[Get sloth]({{< relref "./install.md" >}}) and execute with one of the examples:

```bash
sloth generate -i ./examples/getting-started.yml
```

From the spec, you will obtain the rules for Prometheus with the generated SLO recording rules and alert rules.

{{< tabs "getting-started" >}}
{{< tab "SLO spec" >}}
{{< include file="sync/examples/getting-started.yml" language="yaml" options="linenos=false" >}}
{{< /tab >}}
{{< tab "Generated" >}}
{{< include file="sync/examples/_gen/getting-started.yml" language="yaml" options="linenos=false" >}}
{{< /tab >}}
{{< /tabs >}}
