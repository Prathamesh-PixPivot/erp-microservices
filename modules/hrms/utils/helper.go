package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// ConvertTimestampToISO Convert gRPC Timestamp to ISO 8601 string
func ConvertTimestampToISO(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return time.Unix(ts.Seconds, int64(ts.Nanos)).Format(time.RFC3339)
}

// FormatTimestamp Convert time.Time to string in database format
func FormatTimestamp(ts time.Time) string {
	return ts.UTC().Format("2006-01-02 15:04:05") // ✅ Returns timestamp as "YYYY-MM-DD HH:MM:SS"
}

func GetString(s *string) string {
	if s == nil {
		return "" // Return empty string if nil
	}
	return *s
}
