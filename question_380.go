package leetcode

import (
	"math/rand"
)

type RandomizedSet struct {
	n int // 当前元素个数
	q []int
	m map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		n: 0,
		q: []int{},
		m: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.q = append(this.q, val)
	this.m[val] = this.n
	this.n++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// idx 是空洞
	if idx, ok := this.m[val]; !ok {
		return false
	} else {
		last := this.q[this.n-1] // 获取最后一个元素
		if last == val {         // 最后一个元素被删除的情况
			delete(this.m, val)
			this.q = this.q[:this.n-1]
			this.n--
		} else { // 不是最后一个元素被删除的情况
			delete(this.m, val)
			// 将last元素放置到空洞位置
			this.q[idx] = last
			this.m[last] = idx
			this.q = this.q[:this.n-1]
			this.n--
		}
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.q[rand.Intn(this.n)]
}
