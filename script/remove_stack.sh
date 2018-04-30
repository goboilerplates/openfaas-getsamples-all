#!/bin/sh
faas-cli remove -f function.yml
docker stack rm func