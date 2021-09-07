---
title: "SLI plugins"
weight: 020
---

SLI plugins are small Go plugins (using [Yaegi]) that can be loaded on Sloth start.

These plugins can be referenced as an SLI on the SLO specs and will return a raw SLI type.

## Common plugins

Sloth maintains a library with [common SLI plugins][common-sli-plugins] that can be used on your SLOs or used as examples to develop your own ones.

- Check plugins source code [here][common-sli-plugins].
- Check the documentation [here]({{< relref "../sli-plugins" >}}).

## Developing plugins

### [`prometheus/v1`][plugins-v1]

Developing a [`prometheus/v1`][plugins-v1] SLI plugin is very easy, however you need to meet some requirements:

- The plugin version used as a global called `SLIPluginVersion`.
- A plugin ID global called `SLIPluginID`.
- A Plugin logic function called `SLIPlugin`.
- The plugin must be in a single file named `plugin.go`.
- Plugins only can use the Go standard library (`reflect` and `unsafe` packages can't be used).
- Plugin received options are a `map[string]string` to avoid `interface{}` problems on dynamic execution code, the conversion to specific types are responsibility of the plugin.
- Plugins don't depend on go modules, GOPATH or similar (thanks to the previous requirements).

Sloth knows how to autodiscover plugins giving a path (`--sli-plugins-path`), and will load all the discovered ones.

A very simple example:

from `plugins/x/y/plugin.go`

```go
package testplugin

import "context"

const (
  SLIPluginVersion = "prometheus/v1"
  SLIPluginID      = "test_plugin"
)

func SLIPlugin(ctx context.Context, meta, labels, options map[string]string) (string, error) {
  return "rate(my_raw_error_ratio_query{}[{{.window}}])", nil
}
```

Used in SLO spec:

```yaml
version: "prometheus/v1"
service: "myservice"
slos:
  - name: "some-slo"
    objective: 99.9
    sli:
      plugin:
        id: "test_plugin"
        options:
          opt1: "something"
          opt2: "something"
    alerting:
#...
```

On spec load, Sloth will execute the referenced plugins with the options and use the result as a Raw SLI type, the one that returns the error ratio query.

## Why should I use plugins?

By default you shouldn't unless you have scenarios where they can simplify, add security or improve the SLO adoption on the team/company. Some examples:

- SLI custom validation (parameters, queries...).
- Company wide precreated SLIs for common used libraries.
- Complex Prometheus query SLIs.
- Precreated SLIs for the team or company that normally everyones uses on the SLOs.
- OSS SLI plugins that come with the apps, frameworks or libraries (e.g Kubernetes apiserver, http framework X...).
- The company or the team could have a repository with all the shared plugins and everyone could use them and contribute with new ones.
- Automation power when templates are not enough.

[yaegi]: https://github.com/traefik/yaegi
[common-sli-plugins]: https://github.com/slok/sloth-common-sli-plugins
[plugins-v1]: https://github.com/slok/sloth/blob/main/pkg/prometheus/plugin/v1/v1.go
