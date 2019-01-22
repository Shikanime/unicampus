#!/bin/bash

set -ex

kubectl apply \
  --filename https://github.com/knative/serving/releases/download/v0.3.0/serving.yaml \
  --filename https://github.com/knative/build/releases/download/v0.3.0/release.yaml \
  --filename https://github.com/knative/serving/releases/download/v0.3.0/monitoring.yaml
