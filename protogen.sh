#!/usr/bin/env bash
#
# Generate all protobuf bindings.
# Run from repository root.
set -e
set -u

if ! [[ "$0" =~ "protogen.sh" ]]; then
	echo "must be run from repository root"
	exit 255
fi

if ! [[ $(protoc --version) =~ "3.4.0" ]]; then
	echo "could not find protoc 3.4.0, is it installed + in PATH?"
	exit 255
fi

echo 'Install required packages (clone in GO-like dirs):
github.com/gogo/protobuf
github.com/lyft/protoc-gen-validate
github.com/googleapis/googleapis
github.com/prometheus/client_model
github.com/census-instrumentation/opencensus-proto
'

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/" && pwd -P)"
PROTOBUF_DIR=${PROTOBUF_DIR-${ROOT_DIR}/envoy}

GOGOPROTO_ROOT="${GOPATH}/src/github.com/gogo/protobuf"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"
VALIDATE_PATH="${GOPATH}/src/github.com/lyft/protoc-gen-validate"
GOOGLEAPIS_PATH="${GOPATH}/src/github.com/googleapis/googleapis"
METRICS_PATH="${GOPATH}/src/github.com/prometheus/client_model"
TRACE_PATH="${GOPATH}/src/github.com/census-instrumentation/opencensus-proto/opencensus/proto/trace"

# Envoy protos are designed to be build in a standalone manner (per each file).
# Attempt of build using `*.proto` will result in duplicate errors.
for protoFile in `find ${PROTOBUF_DIR} -type f -name "*.proto"`; do
    echo "Generating ${protoFile}"
	pushd $(dirname "${protoFile}")
		protoc --gogofast_out=plugins=grpc:. \
		    -I=. \
            -I="${GOGOPROTO_PATH}" \
            -I="${ROOT_DIR}" \
            -I="${VALIDATE_PATH}" \
            -I="${GOOGLEAPIS_PATH}" \
            -I="${METRICS_PATH}" \
            -I="${TRACE_PATH}" \
            $(basename "${protoFile}")

		sed -i.bak -E 's/import _ \"gogoproto\"//g' *.pb.go
		sed -i.bak -E 's/import _ \"google\/protobuf\"//g' *.pb.go
		rm -f *.bak
		goimports -w *.pb.go
	popd
done
