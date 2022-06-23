package main

func main() {

}

//Min stack on leetcode
type MinStack struct {
	elem []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.elem = append(this.elem, val)
}

func (this *MinStack) Pop() {
	this.elem = this.elem[:len(this.elem)-1]
}

func (this *MinStack) Top() int {
	return this.elem[len(this.elem)-1]
}

func (this *MinStack) GetMin() int {
	min := int(this.elem[0])
	for i := 0; i < len(this.elem); i++ {
		if this.elem[i] < min {
			min = this.elem[i]
		}
	}
	return min
}
