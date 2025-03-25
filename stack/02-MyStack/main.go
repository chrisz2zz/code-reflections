package main

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Top()
}

type list []int

func (l *list) Pop() int {
	tmp := (*l)[0]
	*l = (*l)[1:]
	return tmp
}

func (l *list) Push(x int) {
	*l = append(*l, x)
}

func (l *list) Top() int {
	return (*l)[0]
}

type MyStack struct {
	length int
	first  list
	second list
}

func Constructor() MyStack {
	return MyStack{
		length: 0,
		first:  make(list, 0),
		second: make(list, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.length++

	for len(this.second) != 0 {
		this.first.Push(this.second.Pop())
	}

	this.first.Push(x)
}

func (this *MyStack) Pop() int {
	this.length--

	for len(this.second) != 0 {
		this.first.Push(this.second.Pop())
	}


	for len(this.first) > 1 {
		this.second.Push(this.first.Pop())
	}

	return this.first.Pop()
}

func (this *MyStack) Top() int {

	for len(this.second) != 0 {
		this.first.Push(this.second.Pop())
	}


	for len(this.first) > 1 {
		this.second.Push(this.first.Pop())
	}

	t := this.first.Pop()

	this.second.Push(t)

	return t
}

func (this *MyStack) Empty() bool {
	if this.length == 0 {
		return true
	}

	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
