---
title:
description: Sloth documentation
geekdocNav: false
geekdocAlign: center
geekdocAnchor: false
---

<img src="media/sloth-logo.svg" alt="sloth logo" width="200" />

Stop using complex specs and processes to create [Prometheus] based [SLOs][google-slo].

_Fast, easy and reliable Prometheus SLO generator._

{{< button size="large" relref="introduction/" >}}Get started{{< /button >}}

{{< columns >}}

## ![](https://icongr.am/fontawesome/star.svg?size=32&color=EC604D) &nbsp; Simple

Lightweight, simple CLI, focused on UX, maintainable SLO specs and a single way of doing things.

<--->

## ![](https://icongr.am/fontawesome/cubes.svg?size=32&color=EC604D)&nbsp; Standards

Based on Google's SRE book framework, SLOs and SLO alerting.

<--->

## ![](https://icongr.am/fontawesome/plug.svg?size=32&color=EC604D) &nbsp; Plugins

Abstracts and extends SLIs using plugins. Removing the need to copy-paste or write complex Prometheus queries.

<--->

## ![](https://icongr.am/material/numeric-1-box.svg?size=32&color=EC604D) &nbsp; One way

Standardizes the SLO implementation across companies and teams by creating a single way of doing SLOs.

{{< /columns >}}

{{< columns >}}

## ![](https://icongr.am/fontawesome/repeat.svg?size=32&color=EC604D) &nbsp; Adaptive

Accepts multiple YAML SLO spec/manifest types declarations, like OpenSLO or Kubernetes CRDs.

<--->

## ![](https://icongr.am/simple/prometheus.svg?size=32&color=EC604D)&nbsp; Prometheus

Delegates and relies on Prometheus stable technology for the metrics and alerts.

<--->

## ![](https://icongr.am/simple/kubernetes.svg?size=32&color=EC604D) &nbsp; Kubernetes

Designed to work with the latest shipping standards and technologies like Gitops and Kubernetes.

<--->

## ![](https://icongr.am/simple/grafana.svg?size=32&color=EC604D) &nbsp; Dashboards

Plug and play Grafana dashboards to observe all the services SLOs status.

{{< /columns >}}

[prometheus]: https://prometheus.io
[google-slo]: https://landing.google.com/sre/workbook/chapters/alerting-on-slos/
