#!/bin/bash
# vim: ai:ts=8:sw=8:noet
set -eufCo pipefail
export SHELLOPTS
IFS=$'\t\n'

command -v git >/dev/null 2>&1 || {
  echo 'please install git'
  exit 1
}

TMP_DIR=./tmp
SLOTH_DIR="${TMP_DIR}/sloth"
SYNC_DIR=./docs-src/sli-plugins
PLUGINS_REPO_DIR="${TMP_DIR}/sli-plugins/plugins"
PLUGINS_DIR="${PLUGINS_REPO_DIR}/plugins"

echo "[*] Cleaning"
rm -rf ${TMP_DIR}
mkdir ${TMP_DIR}
rm -rf ${SYNC_DIR}
mkdir ${SYNC_DIR}

echo "[*] Cloning repositories"
git clone git@github.com:slok/sloth-common-sli-plugins -b main --depth 1 "${PLUGINS_REPO_DIR}"

echo "[*] Syncing SLI plugins"
PREV_DIR="${PWD}"
cd ${PLUGINS_DIR}
find . -name "README.md" | while read fname; do
  # Get group name.
  group_name=$(dirname "${fname}")

  # Remove "./" prefix.
  group_name=${group_name#./}

  # Replace "/" with "-"
  id="${group_name//\//-}"

  # Copy readme.
  cp "${fname}" "${PREV_DIR}/${SYNC_DIR}/${id}.md"
done
