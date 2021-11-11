---
title: "SLO validation"
weight: 030
---

Sloth validates the spec on generation, however, on specific steps of the SLO generation process, we only want to validate a group of SLOs.

For this purpose Sloth comes with a helpful command called `validate`. It will discover all the specs recursively and apply the same generation process as `generate` (including plugins, options...) but discarding the result.

Example that validates all SLOs in a directory (including subdirectories) and excludes all in spec files that match `_gen` in the spec path.

```bash
$ sloth validate --input ./examples --sli-plugins-path ./examples/plugins --fs-exclude _gen

INFO[0000] SLI plugins loaded                            plugins=1 version=dev
INFO[0000] Validation succeeded                          slo-specs=13 version=dev
```

This command is very helpful on Gitops and CI pipelines to have a fast feedback loop, independently of the process you are using for generating the SLOs (Kubernetes controller or CLI).
