---
title: "Go library"
weight: 060
---

!!! note "Version"
    The Go library and SLO plugins are available in Sloth `>= v0.15.0`.

!!! tip ":simple-go: Go docs"
    Check out the go docs at [pkg.go.dev](https://pkg.go.dev/github.com/slok/sloth/pkg/lib)

## Introduction

Sloth has traditionally been used as a CLI application or a Kubernetes operator. However, some use cases require **custom SLO generation**, for example, when existing SLO plugins are not enough, or when the SLO workflow needs tight integration with other systems.

The **Sloth Go library** lets you use Sloth’s core logic directly in your Go applications. It includes everything that the main Sloth app supports, such as:

- SLO and SLI plugins.  
- Multiple SLO specification formats.
- Helpers to export results in different formats.

The library is designed to have a **simple, high-level API**, so you can leverage Sloth’s internal logic without dealing with complex or low-level details.

## Components

The library has two main building blocks:

- **SLO generation**: Generates SLOs and returns a structured in-memory representation (domain model) result.  
- **SLO result storage**: Take the generated domain objects and export them into different formats or backends.

This separation allows you to automate SLO generation, customize the pipeline between generation and storage, or even implement your own storage layer.

## Go Docs

Full Go API reference is available at [https://pkg.go.dev/github.com/slok/sloth/pkg/lib](https://pkg.go.dev/github.com/slok/sloth/pkg/lib)

## Examples

### Live Editor

A proof of concept showing Sloth running fully in the browser using [WASM](https://webassembly.org/).  
This example uses the Go library to generate SLOs live on the client side.

- Live editor: [https://live.sloth.dev/](https://live.sloth.dev/)  
- Source code: [https://github.com/slok/sloth-slo-live-editor](https://github.com/slok/sloth-slo-live-editor)

### REST API

Example of exposing the Go library through a simple HTTP API to generate SLOs remotely.

```go
--8<-- "manual/go-lib-rest-api/main.go"
```

### Remote SLO plugins

Example showing how to load and use remote SLO plugins within your Go application.

```go
--8<-- "manual/go-lib-remote-plugin/main.go"
```
