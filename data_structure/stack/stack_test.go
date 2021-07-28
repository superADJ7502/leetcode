package stack

import "testing"

func TestStack_Len(t *testing.T) {
	stack := CreateStack()
	stack.Push(1)
	stack.Push("test")
	if stack.Len() == 2 {
		t.Log("Pass Stack.Len")
		return
	}
	t.Error("Failed Stack.Len")
}

func TestStack_IsEmpty(t *testing.T) {
	stack := CreateStack()
	if stack.IsEmpty() {
		t.Log("Pass Stack.IsEmpty")
		return
	}
	t.Error("Failed Stack.IsEmpty")
}
