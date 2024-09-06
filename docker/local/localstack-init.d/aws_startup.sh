#!/bin/sh
echo "Init localstack sqs"

awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name DEV_BACKOFFICE_STATUS_COUNT_TEST
