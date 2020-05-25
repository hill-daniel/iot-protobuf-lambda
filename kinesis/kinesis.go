package kinesis

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/protobuf/proto"
	"github.com/hill-daniel/iot-protobuf-lambda"
	pb "github.com/hill-daniel/iot-protobuf-lambda/proto"
	"github.com/pkg/errors"
)

// Handler accepts and processes KinesisEvents
type Handler struct {
	db iot.DeviceBase
}

// NewHandler creates a Lambda handler which accepts and processes KinesisEvents
func NewHandler(db iot.DeviceBase) *Handler {
	return &Handler{db: db}
}

// HandleRequest handles the incoming Kinesis event
func (h *Handler) HandleRequest(ctx context.Context, event events.KinesisEvent) (string, error) {
	var errs []error
	for _, r := range event.Records {
		if err := process(r.Kinesis, h.db); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return "failed", fmt.Errorf("failed to process %d kinesis records out of %d", len(errs), len(event.Records))
	}
	return "ok", nil
}

func process(record events.KinesisRecord, db iot.DeviceBase) error {
	device := &pb.Device{}
	if err := proto.Unmarshal(record.Data, device); err != nil {
		return errors.Wrap(err, "failed to unmarshal protobuf")
	}
	return db.Store(device)
}
