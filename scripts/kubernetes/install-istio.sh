#!/bin/bash

set -ex

curl -L https://github.com/knative/serving/releases/download/v0.2.2/istio.yaml \
  | sed 's/LoadBalancer/NodePort/' \
  | kubectl apply --filename -

kubectl label namespace default istio-injection=enabled
