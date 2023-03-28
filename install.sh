#!/bin/bash

set -e

REPO="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PB=$REPO/protobuf
PROTOC_DIR=${REPO}/.protoc

cd $PB

rm -rf $PROTOC_DIR && mkdir -p $PROTOC_DIR

source $REPO/versions.sh

PROTO_ARCH="linux-x86_64"

if [[ "$OSTYPE" == "darwin"* ]]; then
  PROTO_ARCH="osx-x86_64"
  if [[ $(uname -m) == 'arm64' ]]; then
    PROTO_ARCH="osx-aarch_64"
  fi
fi

wget -nv -O $PROTOC_DIR/protoc-${PROTOC_VERSION}-${PROTO_ARCH}.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTO_ARCH}.zip
unzip -qq -o $PROTOC_DIR/protoc-${PROTOC_VERSION}-${PROTO_ARCH}.zip -d ${PROTOC_DIR}
rm -f $PROTOC_DIR/protoc-${PROTOC_VERSION}-${PROTO_ARCH}.zip $PROTOC_DIR/readme.txt
GOBIN=$PROTOC_DIR/bin go install github.com/twitchtv/twirp/protoc-gen-twirp@${PROTO_TWIRP_VERSION}
GOBIN=$PROTOC_DIR/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTO_GO_VERSION}
GOBIN=$PROTOC_DIR/bin go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@${PROTO_VTPROTO_VERSION}
GOBIN=$PROTOC_DIR/bin go install github.com/yoheimuta/protolint/cmd/protolint@${PROTO_LINT_VERSION}
(
  cd $REPO/tools
  ./install.sh
)

if command -v dart &> /dev/null
then
  PUB_CACHE=${PROTOC_DIR}/.pub-cache dart pub global activate protoc_plugin 20.0.1
  cp ${PROTOC_DIR}/.pub-cache/bin/protoc-gen-dart ${PROTOC_DIR}/bin/protoc-gen-dart
fi

# Swift also requires installing Xcode and other tools which are required for iOS
# development. Without this it will error out complaining that the module 'PackageDescription'
# is not found while building. Hence, providing an option to disable it altogether.
if [[ -z $DISABLE_SWIFT ]]; then
  if command -v swift &> /dev/null
  then
    git clone --depth 1 --branch 1.20.2 https://github.com/apple/swift-protobuf $PROTOC_DIR/.swift-protobuf
    (cd $PROTOC_DIR/.swift-protobuf && swift build -c release)
    ln -s $PROTOC_DIR/.swift-protobuf/.build/release/protoc-gen-swift $PROTOC_DIR/bin/protoc-gen-swift
  fi
else
  echo 'Swift is disabled. Skipping...'
fi

if command -v yarn &> /dev/null
then
  mkdir $PROTOC_DIR/.typescript-protobuf
  (cd $PROTOC_DIR/.typescript-protobuf && yarn add @protobuf-ts/plugin@2.8.1 --no-lockfile --disable-pnp)
  ln -s $PROTOC_DIR/.typescript-protobuf/node_modules/.bin/protoc-gen-ts $PROTOC_DIR/bin/protoc-gen-ts
fi
