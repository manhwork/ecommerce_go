package helper

import "time"

func ParseTimeString(timestring string) string {
	result, _ := time.Parse(time.RFC3339, timestring)

	return result.Format("02/01/2006") // 02/01/2006 15:04:05
}
