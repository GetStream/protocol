#!/bin/bash

TOOLS="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
(
  cd "$TOOLS" || exit 1
  go install ./protoc-gen-dart-twirp
  go install ./protoc-gen-swift-twirp
)
