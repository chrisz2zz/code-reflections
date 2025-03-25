package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1

	for ; l < len(nums); l++ {
		if nums[l] == val {
			break
		}
	}

	for ; r >= 0; r-- {
		if nums[r] != val {
			break
		}
	}

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]

		for ; l < len(nums); l++ {
			if nums[l] == val {
				break
			}
		}

		for ; r >= 0; r-- {
			if nums[r] != val {
				break
			}
		}
	}

	return len(nums[:l])
}
