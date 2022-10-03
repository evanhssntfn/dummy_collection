#! /bin/bash

docker rm -f elasticsearch
docker run -d --name elasticsearch -p 9201:9201 -e discovery.type=single-node \
    -v elasticsearch:/usr/share/elasticsearch/data \
    docker.elastic.co/elasticsearch/elasticsearch:8.4.0
docker ps
