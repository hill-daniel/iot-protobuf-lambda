package pb

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

// ConvertToUnixNanoString converts protobuf timestamp to unix nano second string.
// The default marshaller creates a map instead of a single value
func ConvertToUnixNanoString(timestamp *timestamppb.Timestamp) string {
	ts, err := ptypes.Timestamp(timestamp)
	if err != nil {
		fmt.Printf("faled to parse pb time %v, falling back to now()", timestamp)
		ts = time.Now()
	}
	unixNano := ts.UnixNano()
	return strconv.FormatInt(unixNano, 10)
}
