package sequential

func Search(arr []int, v int) int {
	for idx, val := range arr {
		if val == v {
			return idx
		}
	}
	return -1
}
