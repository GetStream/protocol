#!/bin/bash

# do `source versions.sh` to load vars

# Versions are pinned because when new release of
# these dependencies would be released - there will be
# diff present because of changed versions in generated files.
# It is better to update this values when needed,
# then have unpredictable built definitions.
export PROTOC_VERSION=21.5
export PROTO_TWIRP_VERSION=v8.1.2
export PROTO_GO_VERSION=v1.28.1
export PROTO_VTPROTO_VERSION=v0.3.0
export PROTO_LINT_VERSION=v0.39.0
# Must match the SwiftProtobuf runtime version that the consuming Swift SDKs
# link against (see stream-video-swift Package.swift). Generating with an older
# protoc-gen-swift emits code that trips Swift 6 concurrency checking
# (`static var allCases`, non-Sendable `_StorageClass`) and deprecated APIs
# (`init(serializedData:)`).
export PROTO_SWIFT_VERSION=1.38.1
