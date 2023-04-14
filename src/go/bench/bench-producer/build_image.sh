#!/bin/bash

# Build the docker image
docker build -t kafka-bench-producer -f Dockerfile.producer --cache-from kafka-bench-producer .