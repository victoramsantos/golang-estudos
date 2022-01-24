#!/bin/sh
echo "Init localstack sqs"

# Setup Food Queue
awslocal sqs create-queue --queue-name FoodsQueue-dlq --attributes '{"MessageRetentionPeriod": "259200"}'
awslocal sqs create-queue --queue-name FoodsQueue --attributes '{"RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:000000000000:FoodsQueue-dlq\",\"maxReceiveCount\":\"3\"}", "VisibilityTimeout": "3600"}'

# Setup Food Queue
awslocal sqs create-queue --queue-name DrinksQueue-dlq --attributes '{"MessageRetentionPeriod": "259200"}'
awslocal sqs create-queue --queue-name DrinksQueue --attributes '{"RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:000000000000:DrinksQueue-dlq\",\"maxReceiveCount\":\"3\"}", "VisibilityTimeout": "3600"}'

# Setup BarCounter Queue
awslocal sqs create-queue --queue-name BarCounterQueue-dlq --attributes '{"MessageRetentionPeriod": "259200"}'
awslocal sqs create-queue --queue-name BarCounterQueue --attributes '{"RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:000000000000:BarCounterQueue-dlq\",\"maxReceiveCount\":\"3\"}", "VisibilityTimeout": "3600"}'