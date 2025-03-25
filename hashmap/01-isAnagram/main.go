package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	idx := map[rune]int{}

	for _, v := range s {
		idx[v]++
	}

	for _, v := range t {
		if idx[v] == 0 {
			return false
		}

		idx[v]--
		if idx[v] == 0 {
			delete(idx, v)
		}
	}

	if len(idx) == 0 {
		return true
	}

	return false
}
