#!/bin/bash

# Build the docker image
docker build -t kafka-bench-consumer -f Dockerfile.consumer --cache-from kafka-bench-consumer .