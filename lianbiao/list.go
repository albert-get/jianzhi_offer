package lianbiao

type ListNode struct {
	Val  int
	Next *ListNode
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
func ConstructorJiao(arr1 []int,arr2 []int,arr3 []int) (*ListNode,*ListNode){
	head1:=ListNode{}
	p1 := &head1
	for _, v := range arr1 {
		cur := &ListNode{Val: v}
		p1.Next = cur
		p1 = p1.Next
	}
	head2:=ListNode{}
	p2:=&head2
	for _,v :=range arr2{
		cur := &ListNode{Val: v}
		p2.Next=cur
		p2=p2.Next
	}
	head3:=ListNode{}
	p3:=&head3
	for _,v:=range arr3{
		cur := &ListNode{Val: v}
		p3.Next=cur
		p3=p3.Next
	}
	p1.Next=head3.Next
	p2.Next=head3.Next
	return head1.Next,head2.Next
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