package leetcode

type MyStack struct {
	stack []int
	top   int
	N     int
}

/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	queue := make([]int, 0)
	s := MyStack{stack: queue, top: -1, N: 0}
	return s
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.N++
	this.top++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	value := this.stack[this.top]
	this.stack = this.stack[0 : this.N-1]
	this.N--
	this.top--
	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[this.top]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.N == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := ConstructorMyStack();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
