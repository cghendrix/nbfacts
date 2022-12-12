#!/bin/bash

swag init

# build server
docker buildx build \
--build-arg COMMAND=server \
--build-arg DB_USER=${DB_USER} \
--build-arg DB_PASS=${DB_PASS} \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsserver .

# build worker
docker buildx build \
--build-arg COMMAND=worker \
--build-arg DB_USER=${DB_USER} \
--build-arg DB_PASS=${DB_PASS} \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsworker .

gcloud compute ssh --zone "us-west2-a"  --project "nicklebackfacts" --command "sudo systemctl start konlet-startup"  "nbfactsserver"
