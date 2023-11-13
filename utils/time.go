package utils

import (
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

func FormatTime(t interface{}) string {
	return t.(time.Time).Local().Format(TimeLayout)
}

func Comment(c interface{}) string {
	return c.(interface {
		Comment() string
	}).Comment()
}
