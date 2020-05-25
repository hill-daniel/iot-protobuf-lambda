package pb_test

import (
	"github.com/golang/protobuf/ptypes"
	pb "github.com/hill-daniel/iot-protobuf-lambda/proto"
	"github.com/hill-daniel/iot-protobuf-lambda/test"
	"testing"
	"time"
)

func Test_should_convert_ts_to_unix_time_nano_string(t *testing.T) {
	date := time.Date(2020, 5, 1, 13, 37, 0, 0, time.UTC)
	protoTs, err := ptypes.TimestampProto(date)
	test.Ok(t, err)

	nanoString := pb.ConvertToUnixNanoString(protoTs)

	test.Equals(t, "1588340220000000000", nanoString)
}
