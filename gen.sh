#!/bin/bash

set -eu

# Generate CRDs and open-api schema
cue-gen -paths=common-protos -f=./cue.yaml
cue-gen -paths=common-protos -f=./cue.yaml --crd=true -snake=jwksUri,apiKeys,apiSpecs,includedPaths,jwtHeaders,triggerRules,excludedPaths,mirrorPercent