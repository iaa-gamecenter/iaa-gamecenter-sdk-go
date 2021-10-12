package util

import (
	"strconv"
	"time"
)

var (
	LOC, _ = time.LoadLocation("Asia/Shanghai")
)

func TodayUnix() int64 {
	return time.Now().Unix()
}

func TodayUnixStr() string {
	timeUnix := time.Now().Unix()
	return  strconv.FormatInt(timeUnix,10)
}
