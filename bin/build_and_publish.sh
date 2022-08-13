#!/usr/bin/env bash
set -ex

commit=$(git describe --tags --always)
version=$(cat VERSION)

docker build --build-arg version="$version" -t "ghcr.io/gorginz/georgia-hello-world:$commit" .

docker push "ghcr.io/gorginz/georgia-hello-world:$commit"
