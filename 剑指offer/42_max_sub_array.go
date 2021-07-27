package offer

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= 0 {
			nums[i] += nums[i-1]
		}
		ans = max(ans, nums[i])
	}
	return ans
}
