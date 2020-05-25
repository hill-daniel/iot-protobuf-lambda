package iot

import (
	pb "github.com/hill-daniel/iot-protobuf-lambda/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeStampConverter is a function to convert protobuf timestamp to string
type TimeStampConverter func(timestamp *timestamppb.Timestamp) string

// DeviceBase persists device data
type DeviceBase interface {
	Store(device *pb.Device) error
}
