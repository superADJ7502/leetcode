package offer

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || k >= len(arr) {
		return nil
	}
	left, right := 0, len(arr)-1
	mid := portion(arr, left, right)
	for mid+1 != k {
		if mid+1 < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = portion(arr, left, right)
	}
	return arr[:k]
}

func portion(arr []int, left, right int) int {
	guard := arr[left]
	for left < right {
		for left < right && arr[right] >= guard {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] < guard {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = guard
	return left
}
