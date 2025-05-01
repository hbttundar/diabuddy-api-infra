package helpers

import "time"

// NowUTC returns the current UTC time.
func NowUTC() time.Time {
	return time.Now().UTC()
}

// FormatDate formats a time.Time into YYYY-MM-DD format.
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDate parses a string in YYYY-MM-DD format to time.Time.
func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}
