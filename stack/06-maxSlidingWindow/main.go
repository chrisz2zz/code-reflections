package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	lt := MyQueue{queue: make([]int, 0)}

	res := make([]int, 0)

	for i := 0; i < k; i++ {
		lt.Push(nums[i])
	}

	res = append(res, lt.Front())

	for i := k; i < len(nums); i++ {
		lt.Pop(nums[i-k])
		lt.Push(nums[i])
		res = append(res, lt.Front())
	}

	return res
}
