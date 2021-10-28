#!/usr/bin/env bash
set -euo pipefail

GO_SRC=$GOPATH/src
GO_PROJECT=github.com/pepeunlimited/go-twirp-starter-kit  # rename
GO_PROTOFILES=$GO_PROJECT/api/protobuf                    # `.proto` location
PROTO_SERVICES=$GO_PROTOFILES/services                    # services/`.proto` location
PROTOC_GEN_SH=$GO_PROJECT/scripts/protoc-gen.sh

echo $PROTO_SERVICES

PROTOC_CMD() {
    FILENAMES=$PROTO_SERVICES/*.proto;
    MUTATED=()
    for FILENAME in $FILENAMES; do
        FILENAME_WITHOUT_PATH=$(basename $FILENAME)
        MUTATED+=("services/${FILENAME_WITHOUT_PATH}")
    done
    # echo "${MUTATED[@]}"
    sh $PROTOC_GEN_SH "${MUTATED[@]}"
}

CURRENT_FOLDER=$PWD;cd "$GO_SRC";PROTOC_CMD; cd "$CURRENT_FOLDER";