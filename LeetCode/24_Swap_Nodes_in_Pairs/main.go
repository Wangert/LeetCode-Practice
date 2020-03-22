package main

import "fmt"

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

	test1 := &ListNode{1, nil}
	test1.Next = &ListNode{2, nil}
	test1.Next.Next = &ListNode{3, nil}
	test1.Next.Next.Next = &ListNode{4, nil}

	result := swapPairs(test1)

	for {

		if result == nil {
			break
		}

		fmt.Println(result.Val)
		result = result.Next
	}

}


func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	//定义当前1节点
	node1 := head
	//定义当前2节点
	node2 := head.Next
	//定义前一组交换后的2节点
	preNewNode2 := &ListNode{}

	//结果头指针
	result := node2

	for {

		node1.Next = node2.Next
		node2.Next = node1

		//将前一组交换后的2节点指向当前组交换后的1节点
		preNewNode2.Next = node2
		//改变前一组交换后的2节点，供下次循环使用
		preNewNode2 = node1

		//将node1后移
		node1 = node1.Next

		if node1 == nil {
			break
		}
		//将node2后移
		node2 = node1.Next

		if node2 == nil {
			break
		}

	}

	return result

}