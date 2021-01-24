package main

import "fmt"

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type ListNode struct {
	  Val int
	  Next *ListNode
}




func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	sum := 0
	for l1!=nil && l2!=nil {
		sum = l1.Val + l2.Val
		if sum < 10 {
			l1.Val = sum
			if l1.Next == nil && l2.Next != nil {
				l1.Next = &ListNode{Val: 0,Next: nil}
			}
		}else {
			l1.Val = sum - 10
			if l1.Next == nil {
				l1.Next = &ListNode{Val: sum/10,Next: nil}
			}else {
				l1.Next.Val += sum/10
			}
		}
		if l2.Next == nil && l1.Next != nil{
			l2.Next = &ListNode{Val: 0,Next: nil}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return p
}

func main() {
	var l = &ListNode{Val: 1,Next: nil}
	p := l
	fmt.Println(l.Val,l.Next)

	if l.Next == nil {
		l.Next = &ListNode{Val: 2,Next: nil}
	}
	
	fmt.Println("\n\n")

	for p != nil {
		fmt.Println(p.Val,p.Next)
		p = p.Next
	}
}