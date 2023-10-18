#!/bin/bash

docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t ghcr.io/antin0de/comm-relay:latest --push .
