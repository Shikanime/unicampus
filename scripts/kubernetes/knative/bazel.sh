#!/bin/bash

set -ex

kubectl apply -f https://raw.githubusercontent.com/knative/build-templates/master/bazel/bazel.yaml
