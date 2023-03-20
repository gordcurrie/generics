package post118

// Stack
type Stack[T any] []T

func (s *Stack[T]) Push(i T) {
	*s = append(*s, i)
}

func (s *Stack[T]) Pop() T {
	o := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return o
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) < 1
}
