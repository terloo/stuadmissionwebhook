#!/bin/bash

# build webhook
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stuadmissionwebhook
# build docker image
docker build --no-cache -t stuadmissionwebhook:v1 .
rm -rf stuadmissionwebhook

#docker push ${DOCKER_USER}/admission-webhook-example:v1
