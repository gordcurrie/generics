package pre118

type StringStack []string

func (s *StringStack) Push(i string) {
	*s = append(*s, i)
}

func (s *StringStack) Pop() string {
	o := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return o
}

func (s *StringStack) IsEmpty() bool {
	return len(*s) < 1
}

// IntStack
type IntStack []int

func (s *IntStack) Push(i int) {
	*s = append(*s, i)
}

func (s *IntStack) Pop() int {
	o := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return o
}

func (s *IntStack) IsEmpty() bool {
	return len(*s) < 1
}

// InterfaceStack
type InterfaceStack []interface{}

func (s *InterfaceStack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *InterfaceStack) Pop() interface{} {
	o := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return o
}

func (s *InterfaceStack) IsEmpty() bool {
	return len(*s) < 1
}
