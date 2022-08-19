package alertcore

import "time"

func TimeIsOffDay(time time.Time) bool {
	return time.Weekday() >= 6
}

//CompareArrByPosition arr1>arr2 return 1, arr1<arr2 return 1, arr1=arr2 return 0
func CompareArrByPosition(arr1 []int, arr2 []int) int {
	for index, v1 := range arr1 {
		if len(arr2) <= index || v1 > arr2[index] {
			return 1
		}
		if v1 < arr2[index] {
			return -1
		}
	}
	if len(arr2) > len(arr1) {
		return -1
	}
	return 0
}

func TimeIsEveryday(t time.Time, startTime time.Time, endTime time.Time) bool {
	tArr := []int{t.Hour(), t.Minute(), t.Second()}
	startTimeArr := []int{startTime.Hour(), startTime.Minute(), startTime.Second()}
	endTimeArr := []int{endTime.Hour(), endTime.Minute(), endTime.Second()}
	if CompareArrByPosition(tArr, startTimeArr) == 1 && CompareArrByPosition(tArr, endTimeArr) == -1 {
		return true
	}
	return false
}

func TimeIsBlock(t time.Time, startTime time.Time, endTime time.Time) bool {
	if t.Unix() > startTime.Unix() && t.Unix() < endTime.Unix() {
		return true
	}
	return false
}
