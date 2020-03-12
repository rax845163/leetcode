package q155

import (
	"math"
)

type MinStack struct {
	values []int
	min    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values: []int{},
		min:    math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	val := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	if this.min == val {
		this.resetMin()
	}
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) resetMin() {
	min := math.MaxInt32
	for _, v := range this.values {
		if min > v {
			min = v
		}
	}
	this.min = min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
