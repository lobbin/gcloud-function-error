#!/bin/sh

cd cmd/test
go mod vendor
gcloud functions deploy test-function --entry-point=Test --region=europe-west3 --source . --trigger-http --runtime=go113