package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/hill-daniel/iot-protobuf-lambda"
	pb "github.com/hill-daniel/iot-protobuf-lambda/proto"
	"github.com/pkg/errors"
	"os"
)

// DeviceBase stores device data to a DynamoDB table
type DeviceBase struct {
	*dynamodb.DynamoDB
	timeStampConverter iot.TimeStampConverter
}

// New creates a new DynamoDB device data storage
func New(sess *session.Session, tc iot.TimeStampConverter) *DeviceBase {
	db := dynamodb.New(sess)
	return &DeviceBase{db, tc}
}

// Store persists date from the given device
func (db *DeviceBase) Store(device *pb.Device) error {
	attributeMap, err := toAttributeMap(device, db.timeStampConverter)
	if err != nil {
		return err
	}

	if err = db.store(attributeMap); err != nil {
		return err
	}
	return nil
}

func toAttributeMap(device *pb.Device, timeStampConverter iot.TimeStampConverter) (map[string]*dynamodb.AttributeValue, error) {
	itemMap, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal device %v to dynamo item", device)
	}

	unixTs := timeStampConverter(device.LastUpdated)
	itemMap["last_updated"] = &dynamodb.AttributeValue{N: &unixTs}
	return itemMap, nil
}

func (db *DeviceBase) store(attributeMap map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.PutItemInput{
		Item:      attributeMap,
		TableName: aws.String(os.Getenv("dynamo_device_log_table")),
	}

	if _, err := db.PutItem(input); err != nil {
		return errors.Wrapf(err, "failed to store input %v in DynamoDB", input)
	}
	return nil
}
