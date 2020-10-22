package main

import (
	"fmt"
)

// 链表的长度，不包过头
type Node struct {
	Next *Node
	Data int
}
type LinkList struct {
	Header *Node
}

func NewLinkList() *LinkList {
	return &LinkList{
		Header: &Node{
			Next: nil,
			Data: 0,
		},
	}
}

func (l *LinkList) HasRing(head *Node) bool {
	//var trace = make(map[*Node]struct{}, 0)
	//var next = l.Header.Next
	//fmt.Println("next", l.Header, (l.Header).Next)
	//for next != nil {
	//	trace[next] = struct{}{}
	//	next = next.Next
	//	fmt.Println("dd", trace)
	//	if _, ok := trace[next]; ok {
	//		return true
	//	}
	//}
	//
	//return false

	if head == nil {
		return false
	}
	fmt.Println("head", head, head.Next)
	cache := make(map[*Node]struct{})
	for head.Next != nil {
		if _, ok := cache[head]; ok {
			return true
		}
		cache[head] = struct{}{}
		head = head.Next
		fmt.Println("head", head, head.Next)
	}
	return false

}

func (l *LinkList) Insert(node *Node) {
	var n = node
	nn := &Node{
		Next: n.Next,
		Data: n.Data,
	}
	n.Next = nn
	n = nn
	fmt.Println(n)
}

// 判断链表是否有环
// 输出true
func main() {
	var node1 = new(Node)
	var node2 = new(Node)
	var node3 = new(Node)
	var node4 = new(Node)
	var node5 = new(Node)

	node1.Data = 1
	node2.Data = 2
	node3.Data = 3
	node4.Data = 4
	node5.Data = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	l := NewLinkList()

	fmt.Println("node5", node5.Next, l.Header)
	fmt.Println()
	fmt.Println(l.HasRing(node1))
}

/*
package main

import "fmt"

type ListNode struct {
	Data interface{}
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	var node1 = new(ListNode)
	var node2 = new(ListNode)
	var node3 = new(ListNode)
	var node4 = new(ListNode)
	var node5 = new(ListNode)

	node1.Data = 1
	node2.Data = 2
	node3.Data = 3
	node4.Data = 4
	node5.Data = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	//node4.Next = node5

	hasCycle := HasCycle(node1)

	fmt.Printf("This list has cycle? Yes or No: %v\n", hasCycle)
}
*/
/*
	1 x 2 x
	x x 1 2
	x x 1 x 2 x
	x x x x 1 2 x
	x x x x x x 21
*/
