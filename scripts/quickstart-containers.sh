#!/usr/bin/env bash
#!/bin/bash

# Run docker-compose up -d on all arguments
for arg in "$@"
do
    docker-compose up -d "$arg"
    sleep 5
done
docker-compose ps