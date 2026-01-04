---
title: "K8s transformer plugins"
weight: 046
---

!!! warning "Advanced usage"
    K8s transformer plugins are for specific Kubernetes setups. Most users should use the default Prometheus Operator output.

!!! note "Version"
    K8s transformer plugins are available on Sloth `>=v0.16.0`.

!!! note "Kubernetes controller only"
    This feature only works with `kubernetes` specs mode.

## Introduction

K8s transformer plugins transform generated Prometheus rules into custom Kubernetes objects. By default, Sloth generates Prometheus Operator `PrometheusRule` CRs. With transformer plugins, you can output to:

- ConfigMaps
- VictoriaMetrics VMRule CRs
- Google Managed Prometheus Rules
- ... in the end, any custom K8s resource.

These plugins use Kubernetes' `Unstructured` API to dynamically create any resource type.

## Built-in Plugins

Sloth includes built-in transformer plugins. [Check these](https://github.com/slok/sloth/tree/main/internal/plugin/k8stransform)

## Configuration and usage

!!! tip "Loading external custom plugins"

    Use the common `--plugins-path` (`-p`) flag (same as SLI/SLO plugins) to load custom plugins not buil-in in Sloth. Sloth autodiscovers all plugins in the provided paths.

```bash
sloth generate -p ./plugins/ --k8s-transform-plugin-id="mycompany.com/custom-plugin/v1"
```

## Developing Custom Plugins

### Requirements

- Define `PluginVersion = "prometheus/k8stransform/v1"`
- Define unique `PluginID`
- Implement `NewPlugin()` factory function
- Single file named `plugin.go`
- Only Go standard library + [approved packages](#available-packages)

### Plugin Interface

```go
type Plugin interface {
    TransformK8sObjects(
        ctx context.Context,
        kmeta model.K8sMeta,
        sloResult model.PromSLOGroupResult,
    ) (*K8sObjects, error)
}
```

**Parameters:**

- `kmeta`: Kubernetes metadata (namespace, name, labels, annotations)
- `sloResult`: Generated Prometheus rules for all SLOs
- Returns: List of `unstructured.Unstructured` K8s objects

**Can return multiple objects** to split large outputs.

### Available Packages

- Go standard library (no `reflect`, `unsafe`)
- `k8s.io/apimachinery/pkg/apis/meta/v1/unstructured`
- `github.com/slok/sloth/pkg/common/model`
- `github.com/slok/sloth/pkg/common/utils/k8s`
- `github.com/slok/sloth/pkg/prometheus/plugin/k8stransform/v1`

### Example Plugin

Simplified ConfigMap transformer:

```go
--8<-- "manual/k8s-transformer-plugins/plugin-example.go"
```

### Testing

Use Sloth's testing utilities:

```go
import "github.com/slok/sloth/pkg/prometheus/plugin/k8stransform/v1/testing"
```

Simulates real plugin execution environment and catches errors early.

## Examples

Built-in plugins source code: [github.com/slok/sloth/tree/main/internal/plugin/k8stransform](https://github.com/slok/sloth/tree/main/internal/plugin/k8stransform)

