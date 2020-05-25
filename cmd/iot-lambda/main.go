package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hill-daniel/iot-protobuf-lambda/dynamo"
	"github.com/hill-daniel/iot-protobuf-lambda/kinesis"
	pb "github.com/hill-daniel/iot-protobuf-lambda/proto"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDb := dynamo.New(sess, pb.ConvertToUnixNanoString)
	handler := kinesis.NewHandler(dynamoDb)

	lambda.Start(handler.HandleRequest)
}
