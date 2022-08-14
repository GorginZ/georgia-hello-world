#!/usr/bin/env bash
set -euxo pipefail

die() { echo "$1" >&2; exit "${2:-1}"; }

if [[ "$#" != 1 ]]; then
  echo "usage: $0 <development-environment>"
  exit 245
fi

hash docker || die "docker not found"

env=$1

commit=$(git describe --tags --always)
version=$(cat VERSION)
description=$(cat envs/$env-description.txt)

docker build --build-arg version="$version" --build-arg sha="$commit" --build-arg description="$description" -t "ghcr.io/gorginz/georgia-hello-world:$commit" .

docker push "ghcr.io/gorginz/georgia-hello-world:$commit"
