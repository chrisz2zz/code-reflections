package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcd", 2))
}

func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i+=2*k {
		if i+2*k <= len(s) {
			reverse(ss[i:i+k])
		}else if i+2*k > len(s) && i+k <= len(s) {
			reverse(ss[i:i+k])
		} else {
			reverse(ss[i:len(s)])
		}
	}

	return string(ss)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
