package main

import "fmt"

func main() {
	fmt.Println(isHappy(3))
}

func isHappy(n int) bool {
	idx := make(map[int]struct{})
	sum := 0
	for {
		number := n % 10
		n = n / 10
		sum += number * number
		if n == 0 {
			n = sum
			if _, ok := idx[sum]; ok {
				return false
			}
			idx[sum] = struct{}{}
			if sum == 1 {
				return true
			}
			sum = 0
		}
	}

	return true
}
