#!/usr/bin/env bash


go run cmd/dslist.go -logtostderr=true --kubeconfig ~/.kube/config

#go run cmd/demo.go -kubeconfig=/Users/warren.strange/.kube/config -logtostderr=true