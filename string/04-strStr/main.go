package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

func getNext(next []int, needle string) {
	j := 0
	next[0] = j

	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[j] != needle[i] {
			j = next[j-1]
		}
		if needle[j] == needle[i] {
			j++
		}
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

    next := make([]int, len(needle))
	getNext(next, needle)
	j := 0

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}


