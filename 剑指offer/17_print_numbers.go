package offer

import "math"

func printNumbers(n int) []int {
	max := math.Pow10(n)
	var result []int
	for i := 1; i < int(max); i++ {
		result = append(result, i)
	}
	return result
}
