#!/usr/bin/env sh

set -e
set -x

cleanup() {
	docker rm extract-edge-cluster-contract-grpc-builder
}

trap 'cleanup' EXIT

if [ $# -eq 0 ]; then
	current_directory=$(dirname "$0")
else
	current_directory="$1"
fi

cd "$current_directory"/..

docker build -f docker/Dockerfile.buildGrpcContract -t edge-cluster-contract-grpc-builder .
docker create --name extract-edge-cluster-contract-grpc-builder edge-cluster-contract-grpc-builder
docker cp extract-edge-cluster-contract-grpc-builder:/src/contract/grpc/go ./contract/grpc/
