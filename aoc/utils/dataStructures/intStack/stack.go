package intStack

type Stack struct {
	a []int
}

func (s Stack) Get(e int) int {
	return s.a[e]
}

func (s *Stack) Push(e int) {
	s.a = append(s.a, e)
}

func (s Stack) Top() int {
	return s.a[len(s.a)-1]
}

func (s *Stack) Pop() int {
	tmp := s.Top()
	s.a = s.a[:len(s.a)-1]
	return tmp
}
