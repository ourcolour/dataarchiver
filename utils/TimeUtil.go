package utils

import (
	"github.com/jinzhu/now"
	"github.com/ourcolour/dataarchiver/constants/errs"
	"log"
	"strings"
	"time"
)

const DEFAULT_TIME_FORMAT = "2006-01-02 15:04:05"

func ZeroTime() time.Time {
	return time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
}

func FormatDatetime(datetime time.Time, format string) string {
	// 参数
	if 0 == strings.Compare("", format) {
		format = DEFAULT_TIME_FORMAT
	}
	timeObj := now.New(datetime).Local()
	return timeObj.Format(format)
}
func ParseDatetime(timeString string) (time.Time, error) {
	var (
		result time.Time
		err    error
	)

	// 参数
	if 0 == strings.Compare("", timeString) {
		err = errs.ERR_INVALID_PARAMETERS
		return result, err
	}

	result, err = now.ParseInLocation(time.Local, timeString)

	return result, err
}

func MustParseDatetime(timeString string) time.Time {
	result, err := ParseDatetime(timeString)
	if nil != err {
		log.Panic(err)
	}
	return result
}
