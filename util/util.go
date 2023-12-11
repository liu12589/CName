package util

func GetCloseSet(number int32, arr []int32) int32 {
	for _, v := range arr {
		if number <= v {
			return v
		}
	}
	return 0
}
