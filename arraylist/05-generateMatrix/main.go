package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}

	mid := n / 2

	offset := 1
	c := n / 2
	startx, starty := 0, 0
	start := 1
	for c >= 0 {
		i := startx
		j := starty

		for ; j < n-offset; j++ {
			res[i][j] = start
			start++
		}

		for ; i < n-offset; i++ {
			res[i][j] = start
			start++
		}

		for ; j > starty; j-- {
			res[i][j] = start
			start++
		}

		for ; i > startx; i-- {
			res[i][j] = start
			start++
		}

		offset++
		startx++
		starty++
		c--
	}

	if n%2 != 0 {
		res[mid][mid] = start
	}

	return res
}
