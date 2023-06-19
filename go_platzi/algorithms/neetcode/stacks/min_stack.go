package main

type MinStack struct {
	stack []int
	min   []int
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, min: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min) == 0 || min(val, this.min[len(this.min)-1]) == val {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
