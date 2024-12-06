#!/usr/bin/env bash

set -eo pipefail

go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.0
oapi-codegen -config generator-config.yml ../../openapi/spec.yml > advice_kit_api.go
