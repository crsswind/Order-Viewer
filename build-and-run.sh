#!/bin/bash

docker build ./api -t order-viewer
docker build ./import -t order-viewer-import
docker-compose up
