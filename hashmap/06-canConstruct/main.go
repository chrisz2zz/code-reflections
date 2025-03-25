package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {
    idx := map[rune]int{}

	for _, v := range ransomNote {
		idx[v]++
	}

	for _, v := range magazine {
		if _, ok := idx[v]; ok {
			idx[v]--
			if idx[v] == 0 {
				delete(idx, v)
			}
		}
	}

	if len(idx) == 0 {
		return true
	}

	return false
}