package utils

import "time"

// NowUTC returns the current time in UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

// ParseTimestamp parses an RFC3339 timestamp string
func ParseTimestamp(ts string) (time.Time, error) {
	return time.Parse(time.RFC3339, ts)
}

// FormatTimestamp formats a time as RFC3339
func FormatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}
