package dsa

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	if stack == nil {
		t.Errorf("stack is nil")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("actual %v, expected %v", stack.Size(), 3)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value := stack.Pop()
	if *value != 3 {
		t.Errorf("actual %v, expected %v", *value, 3)
	}
	if stack.Size() != 2 {
		t.Errorf("actual %v, expected %v", stack.Size(), 2)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	value := stack.Peek()
	if *value != 3 {
		t.Errorf("actual %v, expected %v", *value, 3)
	}
	if stack.Size() != 3 {
		t.Errorf("actual %v, expected %v", stack.Size(), 3)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()
	if !stack.IsEmpty() {
		t.Errorf("actual %v, expected %v", stack.IsEmpty(), true)
	}
	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("actual %v, expected %v", stack.IsEmpty(), false)
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("actual %v, expected %v", stack.Size(), 3)
	}
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()
	if stack.Size() != 0 {
		t.Errorf("actual %v, expected %v", stack.Size(), 0)
	}
}

func TestStack_ToSlice(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	slice := stack.ToSlice()
	if len(slice) != 3 {
		t.Errorf("actual %v, expected %v", len(slice), 3)
	}
}

func TestStack_Copy(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	copy := stack.Copy()
	if copy.Size() != 3 {
		t.Errorf("actual %v, expected %v", copy.Size(), 3)
	}
}

func TestStack_Reverse(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Reverse()
	slice := stack.ToSlice()
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Errorf("actual %v, expected %v", slice, []int{1, 2, 3})
	}
}

func TestStack_PopEmpty(t *testing.T) {
	stack := NewStack[int]()
	value := stack.Pop()
	if value != nil {
		t.Errorf("actual %v, expected %v", value, nil)
	}
}

func TestStack_PeekEmpty(t *testing.T) {
	stack := NewStack[int]()
	value := stack.Peek()
	if value != nil {
		t.Errorf("actual %v, expected %v", value, nil)
	}
}
