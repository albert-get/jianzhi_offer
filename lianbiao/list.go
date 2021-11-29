package lianbiao

type ListNode struct {
	Val  int
	Next *ListNode
}
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
type ToST struct {
	Arr []int
	Num int
}

func ConstructorST(data []ToST) *Node {
	var nodeArr []Node
	var child *Node
	for _,item :=range data{
		head:=Node{}
		temp:=&head

		for i,v :=range item.Arr{
			p:=&Node{Val: v}
			if child!=nil&&i==0{
				child.Child=p
			}
			if i==item.Num{
				child=p
			}
			temp.Next=p
			if i!=0{
				p.Prev=temp
			}
			temp=temp.Next
		}
		nodeArr=append(nodeArr,head)
	}
	return nodeArr[0].Next
}

func ConstructorHuan(arr []int, n int) *ListNode {
	head := ListNode{}
	p := &head
	h := &ListNode{}
	for i, v := range arr {

		cur := &ListNode{Val: v}
		if i == n {
			h = cur
		}
		p.Next = cur
		p = p.Next
	}
	p.Next = h

	return head.Next
}

func Constructor(arr []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range arr {
		cur := &ListNode{Val: v}
		p.Next = cur
		p = p.Next
	}
	return head.Next
}
func ConstructorJiao(arr1 []int, arr2 []int, arr3 []int) (*ListNode, *ListNode) {
	head1 := ListNode{}
	p1 := &head1
	for _, v := range arr1 {
		cur := &ListNode{Val: v}
		p1.Next = cur
		p1 = p1.Next
	}
	head2 := ListNode{}
	p2 := &head2
	for _, v := range arr2 {
		cur := &ListNode{Val: v}
		p2.Next = cur
		p2 = p2.Next
	}
	head3 := ListNode{}
	p3 := &head3
	for _, v := range arr3 {
		cur := &ListNode{Val: v}
		p3.Next = cur
		p3 = p3.Next
	}
	p1.Next = head3.Next
	p2.Next = head3.Next
	return head1.Next, head2.Next
}
func ReverseList(head *ListNode) *ListNode {
	prev := ListNode{}.Next
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
