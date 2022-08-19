package alert

import (
	"regexp"
	"strconv"
)

//TimeStringToInt64
func TimeStringToInt64(timeStr string) int64 {
	if timeStr == "" || len(timeStr) == 0 || timeStr == "null" || timeStr == "NULL" {
		return 0
	}
	valueStr := regexp.MustCompile("\\d+(\\.*\\d+)*").FindString(timeStr)
	unitStr := regexp.MustCompile("[a-z]+").FindString(timeStr)
	value, _ := strconv.ParseInt(valueStr, 10, 64)
	switch unitStr {
	case "ms":
		value = value / 1000
		break
	case "s":
		break
	case "m":
		value = value * 1000 * 60
		break
	case "h":
		value = value * 1000 * 60 * 24
		break
	case "d":
		value = value * 1000 * 60 * 60 * 24
		break
	}
	return value
}
