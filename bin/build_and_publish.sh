#!/usr/bin/env bash
set -ex

commit=$(git describe --tags --always)

docker build -t "ghcr.io/gorginz/georgia-hello-world:$commit" .
docker push "ghcr.io/gorginz/georgia-hello-world:$commit"
