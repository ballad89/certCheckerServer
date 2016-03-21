package libhelper

import (
	"fmt"
	"os"
	"time"
)

func NewTime(day, month, year int) (time.Time, error) {

	//const dateFormat = "02-03-2012 SAST"
	//loc, err := time.LoadLocation("Africa/Johannesburg")

	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%d-%d-%dT00:00:00Z00:00", year, month, day))
	//t := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return t, err
}

func GetEnvOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}
