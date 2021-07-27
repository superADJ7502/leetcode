package offer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExistWord(t *testing.T) {
	input := [][][]byte{
		{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		{{'a', 'b'}, {'c', 'd'}},
	}
	output := []string{"ABCCED", "abcd"}
	result := []bool{true, false}
	for i := range input {
		assert.Equal(t, exist(input[i], output[i]), result[i], "come on !")
	}
}

func TestMoveCount(t *testing.T) {
	assert.Equal(t, 4146, movingCount(54, 89, 20), "")
}

func TestPrintNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, printNumbers(1))
}

func TestIsSubStructure(t *testing.T) {
	//assert.Equal(t, isSubStructure(SliceToTree([]int{0, 1, 2, 3}, 1), SliceToTree([]int{0, 3, 1}, 1)), false)
	//assert.Equal(t, isSubStructure(SliceToTree([]int{0, 3, 4, 5, 1, 2}, 1), SliceToTree([]int{0, 4, 1}, 1)), true)
	//assert.Equal(t, isSubStructure(SliceToTree([]int{0, -1, 3, 2, 0}, 1), SliceToTree([]int{0}, 1)), false)
	assert.Equal(t, isSubStructure(SliceToTree([]int{0, 4, 2, 3, 4, 5, 6, 7, 8, 9}, 1), SliceToTree([]int{0, 4, 8, 9}, 1)), false)
}

func TestExchange(t *testing.T) {
	assert.Equal(t, exchange([]int{1, 2, 3, 4}), []int{4, 2, 1, 3})
}

func TestSpiralOrder(t *testing.T) {
	assert.Equal(t, spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	assert.Equal(t, spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}

func TestGetLeastNumbers(t *testing.T) {
	//assert.Equal(t, []int{1, 2}, getLeastNumbers([]int{0, 1, 2, 1}, 1))
	assert.Equal(t, []int{1, 2}, getLeastNumbers([]int{3, 2, 1}, 2))
}

func TestPermutation(t *testing.T) {
	assert.Equal(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, permutation("abc"))
}
