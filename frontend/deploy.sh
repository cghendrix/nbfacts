#!/bin/bash

(cd src; npm run generate)

docker buildx build \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsnginx .

gcloud compute ssh --zone "us-west2-a"  --project "nicklebackfacts" --command "sudo systemctl start konlet-startup"  "nginxnb"

