package stringStack

type Stack struct {
	a []string
}

func (s Stack) Get(e int) string {
	return s.a[e]
}

func (s *Stack) Push(e string) {
	s.a = append(s.a, e)
}

func (s *Stack) Pop() string {
	topIndex := len(s.a) - 1
	top := s.a[topIndex]
	s.a = s.a[:topIndex]
	return top
}

func (s *Stack) Top() string {
	return s.a[len(s.a)-1]
}

func (s Stack) Size() int {
	return len(s.a)
}
