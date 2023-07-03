package utils

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

// TimeToString
// @Description: 时间转字符串
// @param: time
// @return string
func TimeToString(time time.Time) string {
	return carbon.FromStdTime(time).String()
}
