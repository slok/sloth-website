#!/bin/sh

# vim: ai:ts=8:sw=8:noet
set -eufCo pipefail
export SHELLOPTS
IFS=$'\t\n'

command -v mkdocs >/dev/null 2>&1 || {
    echo 'please install mkdocs'
    exit 1
}

mkdocs serve --dev-addr 0.0.0.0:8000
