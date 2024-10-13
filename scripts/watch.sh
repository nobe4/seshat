#!/usr/bin/env bash

find . -type f -name '*.go' \
    | entr -c -r go run main.go
