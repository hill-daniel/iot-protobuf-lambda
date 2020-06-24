# Lambda function for processing protocol buffers messages with AWS IoT Core

## Part of the codecentric blog post ["Processing protocol buffers messages with AWS IoT Core"](https://blog.codecentric.de/en/2020/06/processing-protobufs-with-iot-core)

This is the AWS Lambda code for processing protocol buffers messages and storing contained data into DynamoDB.

## Usage
Execute build.sh or build and zip manually. Afterwards provide the relative path to the root folder of this project to the CDK project and deploy the stack.

## Build (manually)
* `GOOS=linux go build cmd/iot-lambda/main.go` - compile golang code for linux architecture (i.e. AWS Lambda runtime environment)
* `zip function.zip main` - add to zip archive
