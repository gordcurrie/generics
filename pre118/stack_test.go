package pre118_test

import (
	"generics/pre118"
	"testing"
)

func TestStringStack_Push_PoP_Many(t *testing.T) {
	var s pre118.StringStack
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

func TestIntStack_Push_PoP_Many(t *testing.T) {
	var s pre118.IntStack
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

func TestInterfaceStack_Push_PoP_Many(t *testing.T) {
	var s pre118.InterfaceStack
	in := []interface{}{1, "two", false}

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
