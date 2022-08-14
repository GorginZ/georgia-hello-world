#!/usr/bin/env bash
set -exo pipefail

die() { echo "$1" >&2; exit "${2:-1}"; }

hash docker || die "docker not found"

docker compose run --rm lint run -v
