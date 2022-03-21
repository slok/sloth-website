#!/bin/bash
# vim: ai:ts=8:sw=8:noet
set -eufCo pipefail
export SHELLOPTS
IFS=$'\t\n'

command -v docker >/dev/null 2>&1 || {
    echo 'please install docker'
    exit 1
}

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1
[ -z "$DOCKER_FILE_PATH" ] && echo "DOCKER_FILE_PATH env var is required." && exit 1

IMAGE_TAG="${IMAGE}:${VERSION}"

echo "Building image ${IMAGE_TAG}..."
docker build \
    -t "${IMAGE_TAG}" \
    -f "${DOCKER_FILE_PATH}" .
