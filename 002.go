package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	pre := head
	flag := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			flag += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			flag += l2.Val
			l2 = l2.Next
		}
		cur := &ListNode{Val: flag % 10}
		pre.Next = cur
		pre = cur
		flag = flag / 10
	}

	if flag != 0 {
		cur := &ListNode{Val: flag % 10}
		pre.Next = cur
	}

	return head.Next
}

func main()  {
	x1 := &ListNode{Val:5}
	//x2 := &ListNode{Val:4}
	//x3 := &ListNode{Val:3}

	y1 := &ListNode{Val:5}
	//y2 := &ListNode{Val:6}
	//y3 := &ListNode{Val:4}
	//x1.Next = x2
	//x2.Next = x3

	//y1.Next = y2
	//y2.Next = y3

	ans := addTwoNumbers(x1, y1)
	for {
		if ans == nil {
			break
		}
		fmt.Print(ans.Val)
		ans = ans.Next
	}

}
