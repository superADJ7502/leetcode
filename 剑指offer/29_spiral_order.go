package offer

func spiralOrderReverse(arr []int) []int {
	s, e := 0, len(arr)-1
	for s <= e {
		arr[s], arr[e] = arr[e], arr[s]
		s++
		e--
	}
	return arr
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	var result []int
	for left <= right && top <= bottom {
		// left -> right
		if left <= right && top <= bottom {
			result = append(result, matrix[top][left:right+1]...)
			top++
		}

		// top -> bottom
		t := top
		for left <= right && t <= bottom {
			result = append(result, matrix[t][right])
			t++
		}
		right--

		// right -> left
		if left <= right && top <= bottom {
			result = append(result, spiralOrderReverse(matrix[bottom][left:right+1])...)
			bottom--
		}

		//bottom -> top
		b := bottom
		for left <= right && b >= top {
			result = append(result, matrix[b][left])
			b--
		}
		left++
	}
	return result
}
