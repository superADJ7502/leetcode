package offer

func validateStackSequences(pushed []int, popped []int) bool {
	i := 0
	stack := make([]int, 0, len(pushed))
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
