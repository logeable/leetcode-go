package solution

type MinStack struct {
	store []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.store = append(this.store, x)

	currentMin := x
	if len(this.min) > 0 {
		lastMin := this.min[len(this.min)-1]
		if lastMin < currentMin {
			currentMin = lastMin
		}
	}
	this.min = append(this.min, currentMin)
}

func (this *MinStack) Pop() {
	if len(this.store) > 0 {
		this.store = this.store[:len(this.store)-1]
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.store) == 0 {
		return -1
	}
	return this.store[len(this.store)-1]
}

func (this *MinStack) Min() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return -1
}
