package main

func main() {

}

//Stack and Queue
type Stack struct {
	items []int32
}

//adds an item to the end of a slice
func (s *Stack) Push(n int32) {
	s.items = append(s.items, n)
}

//remove the value at the end of a
//slice and returns removed value
func (s *Stack) Pop() int32 {
	position := len(s.items) - 1
	valueAtPosition := s.items[position]
	s.items = s.items[:position]
	return valueAtPosition
}

//returns true if stack is empty, else false
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

//checks the last item of a stack and returns it
func (s *Stack) Peek() int32 {
	value := s.items[len(s.items)-1]
	return value
}

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	// Write your code here
	stack1 := Stack{}
	stack2 := Stack{}
	stack3 := Stack{}
	cumulativeHeight1 := int32(0)
	cumulativeHeight2 := int32(0)
	cumulativeHeight3 := int32(0)
	//fills up stack1 stack2 and stack3
	for i := len(h1) - 1; i >= 0; i-- {
		//0+1
		cumulativeHeight1 += h1[i]
		stack1.Push(cumulativeHeight1)
	}
	for i := len(h2) - 1; i >= 0; i-- {
		cumulativeHeight2 += h2[i]
		stack2.Push(cumulativeHeight2)
	}
	for i := len(h3) - 1; i >= 0; i-- {
		cumulativeHeight3 += h3[i]
		stack3.Push(cumulativeHeight3)
	}
	for true {
		//check if any stack is empty, then return 0 if true
		if stack1.IsEmpty() || stack2.IsEmpty() || stack3.IsEmpty() {
			return 0
		}
		cumulativeHeight1 = stack1.Peek()
		cumulativeHeight2 = stack2.Peek()
		cumulativeHeight3 = stack3.Peek()
		//checks if the three e stacks are equal when we peek
		if cumulativeHeight1 == cumulativeHeight2 && cumulativeHeight2 == cumulativeHeight3 {
			return cumulativeHeight1
		}
		//keep popping if they are not equal
		if cumulativeHeight1 >= cumulativeHeight2 && cumulativeHeight1 >= cumulativeHeight3 {
			stack1.Pop()
		} else if cumulativeHeight2 >= cumulativeHeight3 && cumulativeHeight2 >= cumulativeHeight1 {
			stack2.Pop()
		} else if cumulativeHeight3 >= cumulativeHeight2 && cumulativeHeight3 >= cumulativeHeight1 {
			stack3.Pop()
		}
	}
	return stack1.Peek()
}
