#!/usr/bin/env bash

kubectl delete -f hack/sample/directoryservice.yaml

kubectl create -f hack/sample/directoryservice.yaml
