---
title: "SLO period windows"
weight: 050
---

Sloth tries simplifying SLO generation and maintainability, thats why uses safe defaults and removes a lot of knobs/settings. If you don't need custom SLO period windows is encouraged that you use Sloth's 30 day default one.

## Default SLO period catalog

Sloth can load a catalog of SLO period windows when it starts. By default Sloth comes with some well known SLO period windows loaded:

- 30 day: SRE workbook month SLO period.
- 28 day: Similar to the 30 day but in 4 weeks format.

## Custom SLO period catalog

In case you want to fully customize your SLO period catalog (7 day period, adjust 30 day period alerting windows...). You can use [`AlertWindows/v1`](https://github.com/slok/sloth/tree/main/pkg/prometheus/alertwindows/v1) spec.

Sloth will discover all SLO period window specs, load them on the SLO period catalog so the can be use them, by passing `--slo-period-windows-path`.

Some spec examples:

{{< tabs "" >}}
{{< tab "7 day" >}}
{{< include file="sync/examples/windows/7d.yaml" language="yaml" options="linenos=false" >}}
{{< /tab >}}
{{< tab "Custom 30 day" >}}
{{< include file="sync/examples/windows/custom-30d.yaml" language="yaml" options="linenos=false" >}}
{{< /tab >}}
{{< /tabs >}}

## Selecting default SLO period

Kubernetes and raw specs can's select an SLO period using the spec (because of simplicity), so, sloths fallbacks to its default slo period (`30d`). However, this can be changed using `--default-slo-period`.

Some examples:

- Use default catalog and default 30 day (this would be the same as not specifying anything): `--default-slo-period=30d`.
- Use default catalog with 28 day: `--default-slo-period=28d`.
- Use custom catalog with 7 day: `--default-slo-period=7d --slo-period-windows-path=/my/custom/slo/catalog`.
- Use custom catalog with 45 day: `--default-slo-period=45d --slo-period-windows-path=/my/custom/slo/catalog`.
