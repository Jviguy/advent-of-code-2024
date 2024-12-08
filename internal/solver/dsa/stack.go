package dsa

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() *T {
	if len(s.data) == 0 {
		return nil
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &value
}

func (s *Stack[T]) Peek() *T {
	if len(s.data) == 0 {
		return nil
	}
	value := s.data[len(s.data)-1]
	return &value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = []T{}
}

func (s *Stack[T]) ToSlice() []T {
	return s.data
}

func (s *Stack[T]) Copy() *Stack[T] {
	newStack := NewStack[T]()
	newStack.data = append(newStack.data, s.data...)
	return newStack
}

func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}
