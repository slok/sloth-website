---
title: "Installing Sloth"
---

## Releases

- Using github releases: https://github.com/slok/sloth/releases
- Using [asdf]: https://github.com/slok/asdf-sloth
- Install `sloth-bin` [from the Arch User Repository](https://aur.archlinux.org/packages/sloth-bin)

## Docker images

!!! tip
    All merged PRs on `main` branch have a docker image release.


Official multi arch container images in [ghcr](https://github.com/slok/sloth/pkgs/container/sloth).

```bash
docker pull ghcr.io/slok/sloth
```

## Building from source code

Clone the repository and build:

```bash
git clone git@github.com:slok/sloth.git && \
cd ./sloth && \
make build && \
ls -la ./bin
```

## Kubernetes

- [Helm chart][chart]
- [Raw Kubernetes manifests without common SLI plugins][raw-no-plugins]
- [Raw Kubernetes manifests with common SLI plugins][raw-plugins]
- [Kustomize]

[chart]: https://github.com/slok/sloth/tree/main/deploy/kubernetes/helm
[raw-no-plugins]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/raw/sloth.yaml
[raw-plugins]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/raw/sloth-with-common-plugins.yaml
[kustomize]: https://raw.githubusercontent.com/slok/sloth/main/deploy/kubernetes/kustomization.yaml
[asdf]: https://asdf-vm.com/
