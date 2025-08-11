package testutils

import "time"

// TimeParser parses a time string in RFC3339 format and returns a time.Time object.
func TimeParser(value string) (t time.Time) {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return
}
