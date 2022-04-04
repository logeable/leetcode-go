package solution

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (this *NumArray) Update(index int, val int) {
	this.arr[index] = val

}

func (this *NumArray) SumRange(left int, right int) int {
	ret := 0
	for i := left; i <= right; i++ {
		ret += this.arr[i]
	}
	return ret
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
