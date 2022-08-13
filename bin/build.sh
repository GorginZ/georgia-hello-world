#!/usr/bin/env bash
set -ex
registry="test"
commit=$(git describe --tags --always)

docker build -t "$registry:$commit" . 