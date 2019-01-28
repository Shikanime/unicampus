#!/bin/bash

set -ex

protoc \
  -I/usr/local/include \
  -I. \
  -I$(go env GOPATH)/src \
  -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --gofast_out=plugins=grpc:. \
  ./api/education/v1alpha1/generated.proto
