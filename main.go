package main

import (
	"fmt"
	"leetcode/util"
)

func main() {
	tree := util.SliceToTree([]int{0, 1, 2, 3, 4, 5}, 1)
	fmt.Println(tree)
}
