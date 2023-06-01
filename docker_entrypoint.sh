#!/usr/bin/env bash

set -euo pipefail

BRANCH="main"
if [ $# -eq 3 ]; then
  BRANCH=$(echo "$3" | sed -E 's/.*--branch=([^[:space:]]+).*/\1/')
fi

echo "branch used -> $BRANCH"

git clone https://github.com/GetStream/protocol --branch ${BRANCH} /home/protocol

cd /home/protocol
exec /home/protocol/generate.sh $1 $2
