package main

import "fmt"

func main() {
	fmt.Println(isValid("]"))
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

func isValid(s string) bool {
	st := make(stack, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st.Push(v)
		}

		if v == ')' && st.Pop() != '(' {
			return false
		}

		if v == '}' && st.Pop() != '{' {
			return false
		}

		if v == ']' && st.Pop() != '[' {
			return false
		}
	}

	if len(st) != 0 {
		return false
	}

	return true
}
