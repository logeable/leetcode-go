package solution

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			top := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, top)
		}
	}
	if len(this.s2) == 0 {
		return -1
	}
	top := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return top
}
