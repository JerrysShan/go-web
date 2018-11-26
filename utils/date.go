package utils

import (
	"time"
)

// DateFormat is a tool of format time
func DateFormat(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func StrToDate(str string) (time.Time, error) {
	layout := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeStamp(date time.Time) int64 {
	return date.Unix()
}
