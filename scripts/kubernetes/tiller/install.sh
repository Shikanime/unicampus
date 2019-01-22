#!/bin/bash

set -ex

helm init \
  --service-account tiller
