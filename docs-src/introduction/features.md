---
title: "Features"
---

- Simple, maintainable and understandable SLO spec.
- Reliable SLO metrics and alerts.
- Based on [Google SLO][google-slo] implementation and [multi window multi burn][mwmb] alerts framework.
- Autogenerates Prometheus SLI recording rules in different time windows.
- Autogenerates Prometheus SLO metadata rules.
- Autogenerates Prometheus SLO [multi window multi burn][mwmb] alert rules (Page and warning).
- SLO spec validation (including `validate` command for Gitops and CI).
- Customization of labels, disabling different type of alerts...
- A single way (uniform) of creating SLOs across all different services and teams.
- Automatic [Grafana dashboard](./dashboards.md) to see all your SLOs state.
- Single binary and easy to use CLI.
- Kubernetes ([Prometheus-operator]) support.
- Kubernetes Controller/operator mode with CRDs.
- Support different SLI types.
- Support for [SLI plugins](../usage/plugins.md)
- A library with [common SLI plugins][common-sli-plugins].
- [OpenSLO] support.

[google-slo]: https://landing.google.com/sre/workbook/chapters/alerting-on-slos/
[mwmb]: https://landing.google.com/sre/workbook/chapters/alerting-on-slos/#6-multiwindow-multi-burn-rate-alerts
[prometheus-operator]: https://github.com/prometheus-operator
[common-sli-plugins]: https://github.com/slok/sloth-common-sli-plugins
[openslo]: https://openslo.com/
