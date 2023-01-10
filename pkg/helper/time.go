package helper

import "time"

func BuildTimestamp() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
