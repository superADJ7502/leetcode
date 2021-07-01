package offer

// 小心数字越界
func movingCount(m int, n int, k int) int {
	box := make([][]int, m)
	result := 0
	for i := range box {
		box[i] = make([]int, n)
	}
	box[0][0] = 1
	for row := range box {
		for col := range box[row] {
			if biz(row)+biz(col) > k {
				continue
			}
			if row-1 >= 0 {
				box[row][col] |= box[row-1][col]
			}
			if col-1 >= 0 {
				box[row][col] |= box[row][col-1]
			}
			if box[row][col] > 0 {
				result++
			}
		}
	}
	return result
}

func biz(n int) int {
	var result int
	for n != 0 {
		result += n % 10
		n /= 10
	}
	return result
}
