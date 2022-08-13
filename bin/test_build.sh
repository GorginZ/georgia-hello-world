#!/usr/bin/env bash
set -ex
registry="validate"
commit=$(git describe --tags --always)

docker build -t "$registry:$commit" . 