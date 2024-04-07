// utils/timeUtils.go
package utils

import "time"

func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}
