package utils

import (
	"strconv"
	"time"
)

func StringTimestampToRFC3339(timestampStr string) (string, error) {
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(timestamp, 0)
	rfc3339Time := t.Format(time.RFC3339)
	return rfc3339Time, nil
}

func Int64TimestampToRFC3339(Timestamp int64) (string, error) {
	t := time.Unix(Timestamp, 0)
	rfc3339Time := t.Format(time.RFC3339)
	return rfc3339Time, nil
}
