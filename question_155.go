package leetcode

type MinStack struct {
	stack    []int
	minIndex int
	n        int
}

/** initialize your data structure here. */
func Constructor() *MinStack {
	s := new(MinStack)
	s.stack = make([]int, 0)
	s.minIndex = -1
	s.n = 0
	return s
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.n == 0 {
		this.minIndex = 0
		this.n = 1
	} else {
		if x < this.stack[this.minIndex] {
			this.minIndex = this.n
		}
		this.n++
	}
}

func (this *MinStack) Pop() {
	if this.n == 0 {
		return
	}
	this.stack = this.stack[0 : this.n-1]
	this.n--
	if this.n == 0 {
		this.minIndex = -1
		return
	}
	if this.minIndex == this.n {
		newIndex := 0
		for i := 0; i < this.n; i++ {
			if this.stack[i] < this.stack[newIndex] {
				newIndex = i
			}
		}
		this.minIndex = newIndex
	}
	return
}

func (this *MinStack) Top() int {
	return this.stack[this.n-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.minIndex]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
