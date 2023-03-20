package post118_test

import (
	"generics/post118"
	"testing"
)

func TestStack_Push_PoP_Many(t *testing.T) {
	var s post118.Stack[int]
	in := []int{1, 2, 3}

	for _, v := range in {
		s.Push(v)
	}

	i := len(in) - 1
	for !s.IsEmpty() {
		o := s.Pop()
		if o != in[i] {
			t.Errorf("Expected: %v Does not equal Actual: %v", in[i], o)
		}
		i--
	}
}

func TestStack_Push_PoP_Many_String(t *testing.T) {
	var s post118.Stack[string]
	in := []string{"one", "two", "three"}

	for _, v := range in {
		s.Push(v)
	}

	i := len(in) - 1
	for !s.IsEmpty() {
		o := s.Pop()
		if o != in[i] {
			t.Errorf("Expected: %v Does not equal Actual: %v", in[i], o)
		}
		i--
	}
}
