#!/bin/bash

docker-compose --file localstack/docker-compose.yml down 
docker-compose --file localenv/docker-compose.yml down 
docker-compose --file localobservability/docker-compose.yml down