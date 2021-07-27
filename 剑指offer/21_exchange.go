package offer

func exchange(nums []int) []int {
	s, e := 0, len(nums)-1
	for s < e {
		if nums[s]%2 == 0 {
			nums[e], nums[s] = nums[s], nums[e]
			e--
		} else {
			s++
		}
	}
	return nums
}
