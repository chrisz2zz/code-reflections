package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

type stack []rune

func (s *stack) Push(x rune) {
	*s = append(*s, x)
}

func (s *stack) Pop() rune {
	if len(*s) == 0 {
		return ' '
	}

	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return tmp
}

func (s *stack) Top() rune {
	if len(*s) == 0 {
		return ' '
	}

	tmp := (*s)[len(*s)-1]

	return tmp
}

func (s *stack) String() string {
	ss := ""

	for _, v := range *s {
		ss += string(v)
	}

	return ss
}

func removeDuplicates(s string) string {
	st := make(stack, 0)

	for _, v := range s {
		if len(st) == 0 {
			st.Push(v)
		} else {
			if st.Top() == v {
				st.Pop()
			} else {
				st.Push(v)
			}
		}
	}

	return st.String()
}
