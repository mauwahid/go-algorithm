package binary

func Search(arr []int, v int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if v == guess {
			return mid
		} else if v < guess {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
