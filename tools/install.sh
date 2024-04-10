#!/bin/bash

TOOLS="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROTOC_DIR=${PROTOC_DIR:-"/opt/.protoc"}

(
  cd "$TOOLS" || exit 1
  GOBIN=$PROTOC_DIR/bin go install ./protoc-gen-dart-twirp
  GOBIN=$PROTOC_DIR/bin go install ./protoc-gen-swift-twirp
)
