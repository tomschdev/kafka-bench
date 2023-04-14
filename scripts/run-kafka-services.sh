#!/bin/bash -e

if [ ! -f $(which docker) ];
then
  echo "You need to install docker first"
  exit 1
fi

if [ ! -f $(which docker-compose) ]
then
  echo "You need to install docker-compose first"
  exit 2
fi

./quickstart-containers.sh zookeeper kafka kafka-connect