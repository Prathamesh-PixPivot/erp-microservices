package utils

import (
	"time"
)

// ConvertStringToProtoTimestamp converts a string (RFC3339) to google.protobuf.Timestamp safely
func ConvertStringToTime(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, err // Return zero time if parsing fails
	}

	return parsedTime, nil
}

func ConvertTimeToString(t time.Time) string {
	if t.IsZero() {
		return "" // Return empty string if time is zero
	}
	return t.Format(time.RFC3339)
}