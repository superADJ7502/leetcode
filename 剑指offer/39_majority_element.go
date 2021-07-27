package offer

func majorityElement(nums []int) int {
	master, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			master = nums[i]
			count++
			continue
		}
		if master == nums[i] {
			count++
		} else {
			count--
		}
	}
	return master
}
