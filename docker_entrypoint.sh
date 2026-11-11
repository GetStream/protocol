#!/usr/bin/env bash

set -euo pipefail

BRANCH="main"
if [ $# -eq 3 ]; then
  BRANCH=$(echo "$3" | sed -E 's/.*--branch=([^[:space:]]+).*/\1/')
fi

PROTOCOL_DIR=$(mktemp -d)
git clone https://github.com/GetStream/protocol --branch "${BRANCH}" --depth=1 --single-branch "$PROTOCOL_DIR"

cd "$PROTOCOL_DIR"
exec "$PROTOCOL_DIR/generate.sh" "$1" "$2"
