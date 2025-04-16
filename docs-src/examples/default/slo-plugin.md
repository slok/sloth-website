---
title: "SLO plugin"
---

=== ":octicons-file-code-16: SLO spec"

    ```yaml
    --8<-- "sync/examples/slo-plugin-getting-started.yml"
    ```

=== ":octicons-checklist-16: Generated"

    ```yaml
    --8<-- "sync/examples/_gen/slo-plugin-getting-started.yml"
    ```

=== ":material-console: Execution logs"

    ```bash
    DEBU[0000] Debug level is enabled                        version=dev
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/contrib/grouped_slo/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/contrib/info_labels/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/contrib/rule_intervals/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=github.com/slok/sloth-test-slo-plugins/spec_as_labels/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/alert_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/debug/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/metadata_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/noop/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/sli_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/validate/v1 version=dev window=30d
    INFO[0000] Plugins loaded                                sli-plugins=0 slo-plugins=10 version=dev window=30d
    INFO[0000] SLO period windows loaded                     svc=alert.WindowsRepo version=dev window=30d windows=2
    INFO[0000] Generating from Kubernetes Prometheus spec    version=dev window=30d
    DEBU[0000] Multiwindow-multiburn alerts generated        out=- slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/alert_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/debug/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/metadata_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/noop/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/sli_rules/v1 version=dev window=30d
    DEBU[0000] SLO plugin discovered and loaded              slo-plugin-id=sloth.dev/core/validate/v1 version=dev window=30d
    INFO[0000] Plugins loaded                                sli-plugins=0 slo-plugins=10 version=dev window=30d
    INFO[0000] SLO period windows loaded                     svc=alert.WindowsRepo version=dev window=30d windows=2
    INFO[0000] Generating from Prometheus spec               version=dev window=30d
    DEBU[0000] Multiwindow-multiburn alerts generated        out=- slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 0                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 1                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 2                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 3                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 4                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 5                                      out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    DEBU[0000] Plugin 99                                     out=- plugin=sloth.dev/core/debug/v1 slo=myservice-requests-availability svc=generate.prometheus.Service version=dev window=30d
    INFO[0000] Prometheus rules written                      groups=3 out=- svc=storageio.StdPrometheusGroupedRulesYAMLRepo version=dev window=30d
    ```
