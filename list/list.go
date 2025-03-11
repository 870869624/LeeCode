package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	d := &ListNode{}
	c := d
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			c.Next = list1
			list1 = list1.Next
		} else {
			c.Next = list2
			list2 = list2.Next
		}
		c = c.Next

	}
	if list1 != nil {
		c.Next = list1
	}
	if list2 != nil {
		c.Next = list2
	}
	return d.Next
}
func main() {
	// 创建两个有序链表
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	fmt.Println("List 1:")
	PrintList(l1)

	fmt.Println("List 2:")
	PrintList(l2)

	// 合并两个链表
	merged := MergeTwoLists(l1, l2)

	fmt.Println("Merged List:")
	PrintList(merged)

	reverse(-123)
}

func PrintList(l *ListNode) {
	c := l
	for c != nil {
		fmt.Println(c.Val, "++")
		c = c.Next
	}
}

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	var num int = int(math.Abs(float64(x)))
	temp := num
	count := 1
	for temp >= 10 {
		temp = temp / 10
		count++
	}

	var res int //记录结果的
	flag := count
	//按照权重来计算
	for i := 0; i < flag; i++ {
		count--
		j := num % 10
		num = num / 10
		res = res + j*int(math.Pow10(count))
	}

	if x < 0 {
		return -res
	} else {
		return res
	}
}
