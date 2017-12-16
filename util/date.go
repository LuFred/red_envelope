package util

import (
	"time"
)

//GetTimeMillisecond 返回毫秒时间戳
func GetTimeMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//GetUnix 根据毫秒返回time对象
func GetUnix(millis int64) time.Time {
	return time.Unix(0, millis*1e6)
}