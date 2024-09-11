#!/bin/sh
echo "Init localstack sqs"

awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name OrderEvents
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name PaymentEvents
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name KitchenEvents
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name OrderPreparationEvents
