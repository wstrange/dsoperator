#!/usr/bin/env bash


GOBIN=$(pwd)/bin go install ./cmd/controller-manager



bin/controller-manager --kubeconfig ~/.kube/config --logtostderr -v 0

