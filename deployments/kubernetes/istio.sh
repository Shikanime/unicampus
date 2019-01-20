#!/bin/bash

ISTIO_PATH = ../third_party/istio
ISTIO_CHART_PATH = $ISTIO_PATH/install/kubernetes/helm/

if [ "deploy" = $1 ]; then
  kubectl apply -f $ISTIO_CHART_PATH/helm-service-account.yaml
  helm init --service-account tiller
  helm install $ISTIO_CHART_PATH/istio --name istio --namespace istio-system
elif [ "clean" = $1 ]; then
  kubectl delete -f $ISTIO_CHART_PATH/helm-service-account.yaml
  helm delete --purge istio
  kubectl -n istio-system delete job --all
  kubectl delete -f $ISTIO_CHART_PATH/istio/templates/crds.yaml -n istio-system
fi
