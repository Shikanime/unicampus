#!/bin/bash

if [ "deploy" = $1 ]; then
  kubectl apply -f ./third_party/istio/install/kubernetes/helm/helm-service-account.yaml
  helm install ./third_party/istio/install/kubernetes/helm/istio --name istio --namespace istio-system
elif [ "clean" = $1 ]; then
  kubectl delete -f $ISTIO_CHART_PATH/helm-service-account.yaml
  kubectl -n istio-system delete job --all
fi
