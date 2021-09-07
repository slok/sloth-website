#!/bin/bash
# vim: ai:ts=8:sw=8:noet
set -eufCo pipefail
export SHELLOPTS
IFS=$'\t\n'

command -v git >/dev/null 2>&1 || { echo 'please install git'; exit 1; }

TMP_DIR=./tmp
SYNC_DIR=./sync
SLOTH_DIR="${TMP_DIR}/sloth"
PLUGINS_DIR="${TMP_DIR}/sli-plugins"

echo "[*] Cleaning"
rm -rf ${TMP_DIR}
mkdir ${TMP_DIR}
rm -rf ${SYNC_DIR}
mkdir ${SYNC_DIR}

echo "[*] Cloning repositories"
git clone git://github.com/slok/sloth -b main --depth 1 "${SLOTH_DIR}"
git clone git://github.com/slok/sloth-common-sli-plugins -b main --depth 1 "${PLUGINS_DIR}"

echo "[*] Syncing SLO examples"
cp -r "${SLOTH_DIR}/examples" "${SYNC_DIR}"

