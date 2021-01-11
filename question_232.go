package leetcode

type MyQueue struct {
    Queue []int
    N int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    s := make([]int, 0)
    Q := MyQueue{Queue: s, N: 0}
    return Q
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.Queue = append(this.Queue, x)
    this.N++
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    value := this.Queue[0]
    this.Queue = this.Queue[1:this.N]
    this.N--
    return value
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.Queue[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.N == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

