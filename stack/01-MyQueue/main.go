package main

func main() {

}

type stack []int

func (s *stack) Pop() int {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

func (s *stack) Push(x int) {
	*s = append(*s, x)
}

func (s *stack) Peek() int {
	tmp := (*s)[len(*s)-1]
	return tmp
}

type MyQueue struct {
	first  stack
	second stack
	length int
}

func Constructor() MyQueue {
	return MyQueue{
		first:  make(stack, 0),
		second: make(stack, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.length++

	for len(this.second) != 0 {
		this.first.Push(this.second.Pop())
	}

	this.first.Push(x)
}

func (this *MyQueue) Pop() int {
	this.length--

	for len(this.first) != 0 {
		this.second.Push(this.first.Pop())
	}

	return this.second.Pop()
}

func (this *MyQueue) Peek() int {
	for len(this.first) != 0 {
		this.second.Push(this.first.Pop())
	}

	return this.second.Peek()
}

func (this *MyQueue) Empty() bool {
	if this.length == 0 {
		return true
	}

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
