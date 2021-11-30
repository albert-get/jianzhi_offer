package zhan

type Stack struct {
	nums []interface{}
}

func NewStack() *Stack {
	return &Stack{nums: []interface{}{}}
}

func (s *Stack) Push(n interface{}) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() interface{} {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
type StackInt struct {
	nums []int
}

func NewStackInt() *StackInt {
	return &StackInt{nums: []int{}}
}

func (s *StackInt) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *StackInt) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}
func (s *StackInt) Peek() int {
	res := s.nums[len(s.nums)-1]
	return res
}

func (s *StackInt) Len() int {
	return len(s.nums)
}

func (s *StackInt) IsEmpty() bool {
	return s.Len() == 0
}


