#!/bin/bash

docker-compose --file localstack/docker-compose.yml up -d 
sleep 15
docker-compose --file localenv/docker-compose.yml up -d --build
sleep 15
docker-compose --file localobservability/docker-compose.yml up -d --build 