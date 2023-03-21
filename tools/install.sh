#!/bin/bash

TOOLS="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
REPO=$(dirname $TOOLS)
PROTOC_DIR=${REPO}/.protoc

(
  cd $TOOLS
  GOBIN=$PROTOC_DIR/bin go install ./protoc-gen-dart-twirp
  GOBIN=$PROTOC_DIR/bin go install ./protoc-gen-swift-twirp
)
