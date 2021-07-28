package stack

import "errors"

var (
	ErrEmpty = errors.New("empty value error")
)

type Stack []interface{}

func CreateStack() Stack {
	return Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Cap() int {
	return cap(*s)
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmpty
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element, nil
}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmpty
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil
}
