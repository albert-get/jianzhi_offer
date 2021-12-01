package dui

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Len() int {
	return len(*q)
}

type QueueInt []int

func (q *QueueInt) Push(x int) {
	*q = append(*q, x)
}

func (q *QueueInt) Pop() int {

	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *QueueInt) IsEmpty() bool {
	return len(*q) == 0
}

func (q *QueueInt) Len() int {
	return len(*q)
}
