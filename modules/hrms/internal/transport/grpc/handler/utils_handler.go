package grpc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// toProtoTimestamp safely converts *time.Time to *timestamp.Timestamp
func toProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
