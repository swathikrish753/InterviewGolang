package main

type Stack struct {
	dataStack []int
	maxStack  []int
}

func (s *Stack) Push(value int) {
	s.dataStack = append(s.dataStack, value)

	if len(s.maxStack) == 0 || value >= s.maxStack[len(s.maxStack)-1] {
		s.maxStack = append(s.maxStack, value)
	} else {
		s.maxStack = append(s.maxStack, s.maxStack[len(s.maxStack)-1])
	}
}

func (s *Stack) Pop() int {
	if len(s.dataStack) == 0 {
		panic("Pop from empty stack")
	}
	value := s.dataStack[len(s.dataStack)-1]
	s.dataStack = s.dataStack[:len(s.dataStack)-1]
	s.maxStack = s.maxStack[:len(s.maxStack)-1]
	return value
}

func (s *Stack) Max() int {
	if len(s.maxStack) == 0 {
		panic("Max from empty stack")
	}
	return s.maxStack[len(s.maxStack)-1]
}

func main() {
	s := &Stack{}
	s.Push(3)
	s.Push(5)
	println(s.Max()) // 5
	s.Push(2)
	s.Push(1)
	println(s.Max()) // 5
	s.Pop()
	s.Pop()
	println(s.Max()) // 5
	s.Push(6)
	println(s.Max()) // 6
}
