#!/usr/bin/env bash
set -ex

docker compose run go test -v ./...