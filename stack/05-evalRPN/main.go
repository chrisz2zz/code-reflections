package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

type stack []string

func (s *stack) Push(x string) {
	*s = append(*s, x)
}

func (s *stack) Pop() string {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return tmp
}

func evalRPN(tokens []string) int {
	st := make(stack, 0)

	for _, v := range tokens {
		switch v {
		case "+":
			s1, _ := strconv.ParseInt(st.Pop(), 10, 64)
			s2, _ := strconv.ParseInt(st.Pop(), 10, 64)

			st.Push(strconv.FormatInt(s2+s1, 10))
		case "-":
			s1, _ := strconv.ParseInt(st.Pop(), 10, 64)
			s2, _ := strconv.ParseInt(st.Pop(), 10, 64)

			st.Push(strconv.FormatInt(s2-s1, 10))
		case "*":
			s1, _ := strconv.ParseInt(st.Pop(), 10, 64)
			s2, _ := strconv.ParseInt(st.Pop(), 10, 64)

			st.Push(strconv.FormatInt(s2*s1, 10))

		case "/":
			s1, _ := strconv.ParseInt(st.Pop(), 10, 64)
			s2, _ := strconv.ParseInt(st.Pop(), 10, 64)

			st.Push(strconv.FormatInt(s2/s1, 10))
		default:
			st.Push(v)
		}
	}

	res, _ := strconv.ParseInt(st.Pop(), 10, 64)
	return int(res)
}
