#!/bin/bash

kubectl create secret generic ds --from-file=hack/secrets
