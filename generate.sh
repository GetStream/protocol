#!/bin/bash

set -e

REPO=$PWD
PB=$REPO/protobuf

export PATH=/opt/.protoc/bin:$PWD/.protoc/bin:$PATH

# In some systems like Mac OS 13.0.1, the realpath defaults to
# the system installed one which does not support flags like
# --relative-path. This results in the scripts erroring out.
# The coreutils installed with Homebrew (as mentioned in README)
# installs GNU realpath which could be available as "grealpath"
# which actually works. Hence, the option to customize REALPATH
REALPATH=${REALPATH:-realpath}

GEN_PROFILE=${1:-go}
GEN_OUTPUT=$(${REALPATH} ${2:-${GEN_PROFILE}-sdk})
GEN_GO_IMPORT_PREFIX=${3:-github.com/GetStream/protocol/protobuf}

mkdir -p $GEN_OUTPUT
echo "Generating $GEN_PROFILE SDK in $GEN_OUTPUT"
cd $PB


PROTO_DIRS=$(find $PB -type d)

# All packages that required to generate client SDK (no server-only stuff)
CLIENT_SDK_PROTO_DIRS=$(cat << EOF
$PB/video/sfu/models
$PB/video/sfu/event
$PB/video/sfu/signal_rpc
$PB/video/sfu
EOF
    )

PROTOC_ARGS=""
case $GEN_PROFILE in
  go)
    protoc=$(cat << EOF
    protoc
      -I=$PB
      --go_out=paths=source_relative:$GEN_OUTPUT
      --go-vtproto_out=paths=source_relative,features=marshal+unmarshal+size:$GEN_OUTPUT
      --twirp_out=paths=source_relative:$GEN_OUTPUT
EOF
    )
    # Map imports
    for DIR in $PROTO_DIRS; do
      DIR=$(${REALPATH} --relative-to "$PB" "$DIR")
      FILES=$(find $DIR -maxdepth 1 -name '*.proto' | sort)
      for FILE in $FILES; do
        PROTOC_ARGS="$PROTOC_ARGS --twirp_opt=M${FILE}=${GEN_GO_IMPORT_PREFIX}/${DIR} --go_opt=M${FILE}=${GEN_GO_IMPORT_PREFIX}/${DIR} --go-vtproto_opt=M${FILE}=${GEN_GO_IMPORT_PREFIX}/${DIR}"
      done
    done

    ;;

  dart)
    PROTO_DIRS=$CLIENT_SDK_PROTO_DIRS
    if ! command -v dart &> /dev/null
    then
      echo "dart is required to generate dart profile"
      exit 1
    fi

    protoc=$(cat << EOF
    protoc
      -I=$PB
      --dart_out=$GEN_OUTPUT
      --dart-twirp_out=paths=source_relative:$GEN_OUTPUT
EOF
    )

    # Map imports
    for DIR in $PROTO_DIRS; do
      DIR=$(${REALPATH} --relative-to "$PB" "$DIR")
      FILES=$(find $DIR -maxdepth 1 -name '*.proto' | sort)
      for FILE in $FILES; do
        PROTOC_ARGS="$PROTOC_ARGS --dart-twirp_opt=M${FILE}=${GEN_GO_IMPORT_PREFIX}/${DIR}"
      done
    done

    ;;

  ts)
    PROTO_DIRS=$CLIENT_SDK_PROTO_DIRS
    protoc=$(cat << EOF
    protoc
      -I=$PB
      --ts_out=$GEN_OUTPUT
      --ts_opt long_type_string
      --ts_opt client_generic
      --ts_opt server_none
      --ts_opt eslint_disable
EOF
    )
    ;;

  swift)
    PROTO_DIRS=$CLIENT_SDK_PROTO_DIRS
    if ! command -v swift &> /dev/null
    then
      echo "swift is required to generate swift profile"
      exit 1
    fi

    protoc=$(cat << EOF
    protoc
      -I=$PB
      --swift_opt=FileNaming=FullPath
      --swift_out=$GEN_OUTPUT
      --swift-twirp_out=paths=source_relative:$GEN_OUTPUT
EOF
    )

    # Map imports
    for DIR in $PROTO_DIRS; do
      DIR=$(${REALPATH} --relative-to "$PB" "$DIR")
      FILES=$(find $DIR -maxdepth 1 -name '*.proto' | sort)
      for FILE in $FILES; do
        PROTOC_ARGS="$PROTOC_ARGS --swift-twirp_opt=M${FILE}=${GEN_GO_IMPORT_PREFIX}/${DIR}"
      done
    done

    ;;

  *)
    echo "unknown profile $GEN_PROFILE"
    exit 1
    ;;
esac

for DIR in $PROTO_DIRS; do
  DIR=$(${REALPATH} --relative-to "$PB" "$DIR")
  FILES=$(find $DIR -maxdepth 1 -name '*.proto' | sort)
  ARGS=$PROTOC_ARGS
  if [[ ${FILES} ]]; then
    $protoc ${ARGS} ${FILES}
  fi
done

if [[ $GEN_PROFILE = "go" ]]; then
  # Replace proto.(Un)Marshal with *VT methods
  TWIRP_FILES=$(find "$GEN_OUTPUT" -name '*.twirp.go')
  for FILE in $TWIRP_FILES; do
    if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' -e 's/respBytes, err := proto.Marshal(respContent)/respBytes, err := respContent.MarshalVT()/g' "$FILE"
      sed -i '' -e 's/if err = proto.Unmarshal(buf, reqContent); err != nil {/if err = reqContent.UnmarshalVT(buf); err != nil {/g' "$FILE"
    else
      sed 's/respBytes, err := proto.Marshal(respContent)/respBytes, err := respContent.MarshalVT()/g' "$FILE" > output.txt && mv output.txt "$FILE"
      sed 's/if err = proto.Unmarshal(buf, reqContent); err != nil {/if err = reqContent.UnmarshalVT(buf); err != nil {/g' "$FILE" > output.txt && mv output.txt "$FILE"
    fi
  done
fi
