package brainfuck

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		l := len(*s)
		op := (*s)[l-1]
		*s = (*s)[:l-1]
		return op
	}
}

func (s *Stack) Top() string {
	return (*s)[len(*s)-1]
}
