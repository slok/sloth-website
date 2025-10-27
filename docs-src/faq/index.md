---
title: "F.A.Q"
---

## Why Sloth

Creating Prometheus rules for SLI/SLO framework is hard, error prone and is pure toil.

Sloth abstracts this task, and we also gain:

- Read friendlyness: Easy to read and declare SLI/SLOs.
- Gitops: Easy to integrate with CI flows like validation, checks...
- Reliability and testing: Generated prometheus rules are already known that work, no need the creation of tests.
- Centralize features and error fixes: An update in Sloth would be applied to all the SLOs managed/generated with it.
- Standardize the metrics: Same conventions, automatic dashboards...
- Rollout future features for free with the same specs: e.g automatic report creation.

## SLI?

[Service level indicator][sli]. Is a way of quantify how your service should be responding to user.

TL;DR: What is good/bad service for your users. E.g:

- Requests >=500 considered errors.
- Requests >200ms considered errors.
- Process executions with exit code >0 considered errors.

Normally is measured using events: `good/bad-events / total-events`.

## SLO?

[Service level objective][slo]. A percent that will tell how many [SLI] errors your service can have in a specific period of time.

## Error budget?

An error budget is the ammount of errors (driven by the [SLI]) you can have in a specific period of time, this is driven by the [SLO].

Lets see an example:

- SLI Error: Requests status code >= 500
- Period: 30 days
- SLO: 99.9%
- Error budget: 0.0999 (100-99.9)
- Total requests in 30 days: 10000
- Available error requests: 9.99 (10000 \* 0.0999 / 100)

If we have more than 9.99 request response with >=500 status code, we would be burning more error budget than the available, if we have less errors, we would end without spending all the error budget.

## Burn rate?

The speed you are consuming your error budget. This is key for [SLO] based alerting (Sloth will create all these alerts), because depending on the speed you are consuming your error budget, it will trigger your alerts.

Speed/rate examples:

- 1: You are consuming 100% of the error budget in the expected period (e.g if 30d period, then 30 days).
- 2: You are consuming 200% of the error budget in the expected period (e.g if 30d period, then 15 days).
- 60: You are consuming 6000% of the error budget in the expected period (e.g if 30d period, then 12h hour).
- 1080: You are consuming 108000% of the error budget in the expected period (e.g if 30d period, then 40 minute).

## SLO based alerting?

With SLO based alerting you will get better alerting to a regular alerting system, because:

- Alerts on symptoms ([SLI]s), not causes.
- Trigger at different levels (warning/ticket and critical/page).
- Takes into account time and quantity, this is: speed of errors and number of errors on specific time.

The result of these is:

- Correct time to trigger alerts (important == fast, not so important == slow).
- Reduce alert fatigue.
- Reduce false positives and negatives.

## What are ticket and page alerts?

[MWMB] type alerting is based on two kinds of alerts, `ticket` and `page`:

- `page`: Are critical alerts that normally are used to _wake up_, notify on important channels, trigger oncall...
- `ticket`: The warning alerts that normally open tickets, post messages on non-important Slack channels...

These are triggered in different ways, `page` alerts are triggered faster but require faster error budget burn rate, on the other side, `ticket` alerts
are triggered slower and require a lower and constant error budget burn rate.

## Can I disable alerts?

Yes, use `disable: true` on `page` and `ticket`.

## Grafana dashboard?

Check [grafana-dashboard](../introduction/dashboards.md), this dashboard will load the SLOs automatically.

## CLI VS K8s controller?

If you don't have Kubernetes and you need raw prometheus rules, its easy, the CLI (`generate`) mode is the only one that supports raw prometheus rules.

On the other side if you have Kubernetes (and most likely prometheus-operator). Using [`sloth.slok.dev/v1/PrometheusServiceLevel`](https://raw.githubusercontent.com/slok/sloth/main/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml) CRD will output the same result used as a CLI or used as a Kubernetes controller.

The only difference between the two modes is how Sloth application loads the SLOs manifest. For either mode, the output will be a Prometheus Operator Rules CRD.

Both have pros and cons:

- The CLI in an advanced gitops flow gives you faster feedback loops because of the generation on the CI.
- Using as a controller the CRD integrates better in helm charts and similar because it removes that generation extra step.
- Having the SLO as CRs in K8s, improves the discovery as you can always do `kubectl get slos --all-namespaces`.
- The CLI doesn't require an app running, Sloth CRDs registered... the SLO generation process is simpler, so you have less PoFs.

In a few words, theres no right or wrong answer, pick your own flavour based on your use case: teams size, engineers in the company or development flow...

## SLI types on manifests

`prometheus/v1` (regular) and `sloth.slok.dev/v1/PrometheusServiceLevel` (Kubernetes CRD), support 3 ways of setting SLIs:

- Events: This are based on 2 queries, the one that returns the total/valid number of events and the one that returns the bad events. Sloht will make a query dividing them to get the final error ratio (0-1).
- Raw: This is a single raw prometheus query that when executed will return the error ratio (0-1).
- Plugins: Check [plugins](../sli-plugins/coredns-availability.md) for more information. It reference plugins that will be preloaded and already developed. Sloth will execute them on generation and it will return a raw query. This is the best way to abstract queries from users or having SLOs at scale.

## SLO plugins VS SLI plugins, what should I use?

It depends on your use case:

If you only need to act on the SLI query level — for example, to standardize or slightly modify the queries — then SLI plugins are likely enough. They’re simpler and focused solely on how the SLI is calculated.

If you need more flexibility, such as mutating the final generated Prometheus rules, injecting metadata, rewriting a full prom rule, adding a new Prometheus rule, or making changes before Sloth's validation step, then you should use SLO plugins.

SLO plugins are more powerful and give you full control over the SLO Sloth rules generation process, but they come with added complexity. Use them when you need deep customization.

## SLO plugins at SLO group level or SLO level?

If all the SLOs in a group need the same plugin, use the SLO group level. If you need customization at the most isolated level of SLO, use the SLO level plugins.

Remember plugins are a chain, so you can combine both, set the shared plugin ones at SLO Group level, and specific ones at SLO level.

## Can I override the full Sloth SLO plugin logic?

Yes you can disable the default app plugins and set your own app plugins. These level of SLO plugins will target all SLOs. 

Lets run an example and go step by step:

```bash
sloth generate \
    --disable-default-slo-plugins \ # (1)!
    -i ./examples/getting-started.yml \
    -p /path/to/our/custom/slo/plugins/ \ # (2)!
    -s '{"id": "sloth.dev/core/sli_rules/v1" }' \ # (3)!
    -s '{"id": "my.custom.plugin/rule_intervals/v1", "config": {"interval": {"default": "42m"}}}' # (4)!
```

1. First of all we disable the defualt logic from Sloth (Default app plugins).
    we do this because we want to override all the logic
2. We point the directory for our custom plugins.
3. We want some of the default behaviour from sloth, so we load that specific default plugin, in this case is the SLI rule generation.
4. We load our custom rule interval setter plugin with some config.

With this command example we changed how Sloth works to make multiple things:

- Disabled validation
- Only generate SLI rules
- Add a custom rule interval based on our custom logic.

## Can I set a custom SLO validation logic?

Yes you can, by disabling the default plugins to remove the validation plugin, then setting our own validation and maintaining the default generation plugins from sloth. Example:

```bash
sloth generate \
    -i ./examples/getting-started.yml \
    -p /path/to/our/custom/slo/plugins/ \
    --disable-default-slo-plugins \
    -s '{"id": "my.custom.plugin/custom_validation/v1"}'
    -s '{"id": "sloth.dev/core/sli_rules/v1" }'
    -s '{"id": "sloth.dev/core/metadata_rules/v1" }'
    -s '{"id": "sloth.dev/core/alert_rules/v1" }'
```

## Can the SLO plugins do advanced logic?

Yes, the best way is using an example. This plugin optimizes the performance of the SLI error total window. To do so, the plugins does 2 things:

- Infers metrics from the user defined SLIs by parsing recursively the prometheus expression nodes.
- Creates new rule group with an specific interval

```go
--8<-- "manual/slo-plugin-advanced-faq/plugin.go"
```

## Custom SLO time windows?

Please Check [SLO period windows section](../usage/slo-period-windows.md)

[mwmb]: https://landing.google.com/sre/workbook/chapters/alerting-on-slos/#6-multiwindow-multi-burn-rate-alerts
[sli]: https://landing.google.com/sre/sre-book/chapters/service-level-objectives/#indicators-o8seIAcZ
[slo]: https://landing.google.com/sre/sre-book/chapters/service-level-objectives/#objectives-g0s1tdcz
