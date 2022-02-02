package data

import (
	"strconv"
	"time"
)

var Clock = time.Now()

func RefreshClock() {
	Clock = time.Now()
}

func GetTimestampAsString(timestamp int64) string {
	return strconv.FormatInt(timestamp, 10)
}
