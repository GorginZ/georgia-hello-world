#!/usr/bin/env bash
set -ex

docker compose run --rm lint run -v
