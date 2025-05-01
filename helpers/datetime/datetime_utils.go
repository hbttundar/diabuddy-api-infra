package datetimeutil

import (
	"math/rand"
	"time"
)

func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// NowWithFormat returns the current time using the provided format, if available.
// If no format or an empty format is provided, it defaults to time.RFC3339.
func NowWithFormat(format string) string {
	if format == "" {
		format = time.RFC3339
	}
	return time.Now().UTC().Format(format)
}

// Convert the given time.Time object to default dateTime format.
// If no format is specified, it defaults to time.RFC3339.
func Convert(datetime time.Time) string {
	return datetime.UTC().Format(time.RFC3339)
}

func ConvertWithFormat(datetime time.Time, format string) string {
	if format == "" {
		format = time.RFC3339
	}
	return datetime.UTC().Format(format)
}

// Parse converts a datetime string to a time.Time object using the provided format.
// If no format is provided, it defaults to time.RFC3339.
func Parse(datetimeStr string, format string) (time.Time, error) {
	if format == "" {
		format = time.RFC3339
	}

	return time.ParseInLocation(format, datetimeStr, time.UTC)
}

// DateBetween generates a random date between two specified dates.
func DateBetween(start, end time.Time) time.Time {
	// Ensure start is before an end
	if end.Before(start) {
		start, end = end, start
	}

	// Calculate the difference between start and end
	duration := end.Sub(start)
	randomOffset := time.Duration(rand.Int63n(duration.Nanoseconds()))
	return start.Add(randomOffset)
}
