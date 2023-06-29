#!/bin/bash

set -eu

# Generate all protos
buf generate \
  --path core \
  --path policy

# Generate CRDs and open-api schema
cue-gen -paths=common-protos -f=./cue.yaml
cue-gen -paths=common-protos -f=./cue.yaml --crd=true -snake=jwksUri,apiKeys,apiSpecs,includedPaths,jwtHeaders,triggerRules,excludedPaths,mirrorPercent