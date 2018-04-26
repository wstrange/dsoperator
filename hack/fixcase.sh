#!/usr/bin/env bash

kubebuilder generate

find ./pkg -name \*.go -exec sed -i -e 's|github.com/forgerock|github.com/ForgeRock|' {} +
