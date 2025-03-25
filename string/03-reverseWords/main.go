package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("  hello world  "))
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")


	tmp := make([]string, 0, len(ss))

	for i := 0; i < len(ss); i++ {
		if ss[i] == "" {
			continue
		}

		tmp = append(tmp, ss[i])
	}

	l, r := 0, len(tmp)-1

	for l < r {
		tmp[l],tmp[r] = tmp[r], tmp[l]
		l++
		r--
	}

	return strings.Join(tmp, " ")
}
