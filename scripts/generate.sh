#!/bin/sh
# vim: ai:ts=8:sw=8:noet
set -eufCo pipefail
export SHELLOPTS
IFS=$'\t\n'

command -v mkdocs >/dev/null 2>&1 || {
    echo 'please install mkdocs'
    exit 1
}

# Clean.
rm -rf ./docs && mkdir ./docs

# Generate.
mkdocs build --strict --clean --site-dir docs

# Create Github's custom CNAME.
echo "sloth.dev" >./docs/CNAME
