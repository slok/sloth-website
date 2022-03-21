---
title: "Getting started"
---

[Get sloth](./install.md) and execute with one of the examples:

```bash
sloth generate -i ./examples/getting-started.yml
```

From the spec, you will obtain the rules for Prometheus with the generated SLO recording rules and alert rules.

=== ":octicons-file-code-16: SLO spec"

    ```yaml
    --8<-- "sync/examples/getting-started.yml"
    ```

=== ":octicons-checklist-16: Generated"

    ```yaml
    --8<-- "sync/examples/_gen/getting-started.yml"
    ```
