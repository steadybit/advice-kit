#!/usr/bin/env bash

set -eo pipefail

go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.0
# Skip first line of output as it contains a warning from oapi-codegen.
# It can be ignored, as only supported features are used in the spec.
oapi-codegen -config generator-config.yml ../../openapi/spec.yml | tail -n +2 > advice_kit_api.go
