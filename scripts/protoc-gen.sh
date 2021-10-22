#!/usr/bin/env bash

set -euo pipefail

# WARNING: requires GOPATH 
# vim ~/.bash_profile
# export GOPATH=$HOME/go
GO_SRC=$GOPATH/src
GO_PROJECT=github.com/pepeunlimited/go-twirp-starter-kit # replace with your project name
GO_PROTOFILES=$GO_PROJECT/api/protobuf                   # your `.proto` location
GO_PROTOGEN_OUT=.   # option go_package "foo/bar/pkg/api/v1/resources" defines `.pb.go` `.twirp.go` output location
GO_TWIRP_OUT=.      # option go_package "foo/bar/pkg/api/v1/resources" defines `.pb.go` `.twirp.go` output location
PROTO_PATH=$GO_PROTOFILES
# actual command for protoc
PROTOC_CMD() {
    FILENAMES=("$@")
    # ARGUMENTS:
    # - $ sh protoc-gen.sh enums/foo.proto
    # - $ sh protoc-gen.sh services/bar.proto
    # - $ sh protoc-gen.sh resources/zaa.proto
    echo "protoc-go execute command --proto_path=$PROTO_PATH proto_filenames:"
    echo "${FILENAMES[@]}" | tr ' ' '\n'
    protoc --proto_path=$PROTO_PATH --twirp_out=$GO_TWIRP_OUT --go_out=$GO_PROTOGEN_OUT "${FILENAMES[@]}"
}

if [ "$#" -eq 0 ]; then
    echo "missing argument: FILENAME"
else
    # arguments as array
    ARGS=${1:+"$@"}
    MUTATED=()
    for ARG in $ARGS; do
      MUTATED+=("$PROTO_PATH/${ARG}")
    done
    # execute the cmd
    CURRENT_FOLDER=$PWD;cd "$GO_SRC";PROTOC_CMD "${MUTATED[@]}";cd "$CURRENT_FOLDER";
fi