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
SYNC_DIR=./docs-src/snippets/sync
SYNC_SLO_PLUGINS_DIR_CORE=./docs-src/slo-plugins/core
SYNC_SLO_PLUGINS_DIR_CONTRIB=./docs-src/slo-plugins/contrib
SLOTH_DIR="${TMP_DIR}/sloth"

echo "[*] Cleaning"
rm -rf ${TMP_DIR}
mkdir -p ${TMP_DIR}
rm -rf ${SYNC_DIR}
mkdir -p ${SYNC_DIR}
rm -rf ${SYNC_SLO_PLUGINS_DIR_CORE}
mkdir -p ${SYNC_SLO_PLUGINS_DIR_CORE}
rm -rf ${SYNC_SLO_PLUGINS_DIR_CONTRIB}
mkdir -p ${SYNC_SLO_PLUGINS_DIR_CONTRIB}

echo "[*] Cloning repositories"
git clone git@github.com:slok/sloth -b main --depth 1 "${SLOTH_DIR}"

echo "[*] Syncing SLO examples"
cp -r "${SLOTH_DIR}/examples" "${SYNC_DIR}"

echo "[*] Syncing core SLO plugins"
PREV_DIR="${PWD}"
cd "${SLOTH_DIR}/internal/plugin/slo/core"
find . -name "README.md" | while read fname; do
    # Use the dri as the ID
    id=$(dirname "${fname}")

    # Remove "./" prefix.
    id=${id#./}

    # Replace "/" with "-"
    id="${id//\//-}"

    # Copy readme.
    cp "${fname}" "${PREV_DIR}/${SYNC_SLO_PLUGINS_DIR_CORE}/${id}.md"
done
cd ${PREV_DIR}

echo "[*] Syncing contrib SLO plugins"
cd "${SLOTH_DIR}/internal/plugin/slo/contrib"

find . -name "README.md" | while read fname; do
    # Use the dri as the ID
    id=$(dirname "${fname}")

    # Remove "./" prefix.
    id=${id#./}

    # Replace "/" with "-"
    id="${id//\//-}"

    # Copy readme.
    cp "${fname}" "${PREV_DIR}/${SYNC_SLO_PLUGINS_DIR_CONTRIB}/${id}.md"
done
cd ${PREV_DIR}
