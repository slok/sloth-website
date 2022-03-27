---
title: "Installing Sloth"
---

## :material-tag: Releases

- Using github releases: https://github.com/slok/sloth/releases
- Using [asdf]: https://github.com/slok/asdf-sloth
- Using [Arch linux][arch] AUR: https://aur.archlinux.org/packages/sloth-bin)

## :material-docker: Docker images

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

## :material-kubernetes: Kubernetes

=== "Raw"

    - [Raw Kubernetes manifests without common SLI plugins][raw-no-plugins]
    - [Raw Kubernetes manifests with common SLI plugins][raw-plugins]

=== "Kustomize"

    ```yaml
    apiVersion: kustomize.config.k8s.io/v1beta1
    kind: Kustomization

    resources:
    - https://raw.githubusercontent.com/slok/sloth/v0.10.0/pkg/kubernetes/gen/crd/sloth.slok.dev_prometheusservicelevels.yaml
    - https://raw.githubusercontent.com/slok/sloth/v0.10.0/deploy/kubernetes/raw/sloth-with-common-plugins.yaml
    ```

=== "Helm"

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
