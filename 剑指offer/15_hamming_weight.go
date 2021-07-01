package offer

// num &= num - 1 将二进制最低位的1变成0
func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		num &= num - 1
		ans++
	}
	return ans
}
