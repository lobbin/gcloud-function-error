#!/bin/sh

cd cmd/test_non_internal
go mod vendor
gcloud functions deploy test-function-non-internal --entry-point=Test --region=europe-west3 --source . --trigger-http --runtime=go113