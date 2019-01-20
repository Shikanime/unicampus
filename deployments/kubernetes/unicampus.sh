#!/bin/bash

if [ "deploy" = $1 ]; then
  helm install ./deployments/helm/unicampus --name unicampus --namespace default
elif [ "clean" = $1 ]; then
  helm delete --purge unicampus
fi
