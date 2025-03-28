---
title: "Installing Sloth"
---

## :material-tag: Binary releases

- :material-github: Using github releases: <https://github.com/slok/sloth/releases>
- :material-console-line: Using [asdf]: <https://github.com/slok/asdf-sloth>
- :material-arch: Using [Arch linux][arch] AUR: <https://aur.archlinux.org/packages/sloth-bin>

## :simple-docker: Docker images

!!! tip
    All merged PRs on `main` branch have a docker image release.


Official multi arch container images in [ghcr](https://github.com/slok/sloth/pkgs/container/sloth).

```bash
docker pull ghcr.io/slok/sloth
```

## :material-source-branch: Building from source code

Clone the repository and build:

```bash
git clone git@github.com:slok/sloth.git && \
cd ./sloth && \
make build && \
ls -la ./bin
```

## :simple-kubernetes: Kubernetes

!!! note "CRDs"
    Kubernetes installation will require Sloth CRDs to work correctly.

=== ":material-file: Raw"

    - [Raw Kubernetes manifests without common SLI plugins][raw-no-plugins]
    - [Raw Kubernetes manifests with common SLI plugins][raw-plugins]
    - [CRDs]

=== ":simple-keras: Kustomize"

    ```yaml
    apiVersion: kustomize.config.k8s.io/v1beta1
    kind: Kustomization

    resources:
    - https://raw.githubusercontent.com/slok/sloth/v0.12.0/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
    - https://raw.githubusercontent.com/slok/sloth/v0.12.0/deploy/kubernetes/raw/sloth-with-common-plugins.yaml
    ```

=== ":simple-helm: Helm"

    Helm directory is [here][helm-chart], however, Sloth has helm releases that can be used, example:

    ```bash
    helm repo add sloth https://slok.github.io/sloth
    helm repo update
    helm template sloth/sloth --include-crds
    ```

[chart]: https://github.com/slok/sloth/tree/main/deploy/kubernetes/helm
[raw-no-plugins]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/raw/sloth.yaml
[raw-plugins]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/raw/sloth-with-common-plugins.yaml
[kustomize]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/kustomization.yaml
[asdf]: https://asdf-vm.com/
[arch]: https://archlinux.org/
[helm-chart]: https://github.com/slok/sloth/tree/main/deploy/kubernetes/helm
[CRDs]: https://raw.githubusercontent.com/slok/sloth/v0.12.0/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
