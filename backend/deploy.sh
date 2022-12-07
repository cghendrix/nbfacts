#!/bin/bash

swag init

docker buildx build \
--build-arg DB_USER=${DB_USER} \
--build-arg DB_PASS=${DB_PASS} \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsserver .

gcloud compute ssh --zone "us-west2-a"  --project "nicklebackfacts" --command "sudo systemctl start konlet-startup"  "nbfactsserver"
