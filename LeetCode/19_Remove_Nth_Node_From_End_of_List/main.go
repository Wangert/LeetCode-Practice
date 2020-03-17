package main

import "fmt"

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	result := removeNthFromEnd(head, 2)

	for {

		fmt.Println(result.Val)

		if result.Next != nil {
			result = result.Next
		} else {
			break
		}
	}
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	//用于遍历链表
	tempListNode := head
	//用于存储相应位置节点的地址
	addressMap := make(map[int]*ListNode)
	//节点索引
	index := 0
	//存储结果
	resultListNode := &ListNode{}

	for {
		//记录节点地址
		addressMap[index] = tempListNode

		if tempListNode.Next != nil {
			index++
			tempListNode = tempListNode.Next
		} else {
			break
		}

	}

	//倒数第n个节点的位置索引
	nEndIndex := index - n + 1

	if nEndIndex < 0 {
		return nil
	}

	if nEndIndex == 0 {
		resultListNode = head.Next
	} else if nEndIndex == index {
		addressMap[index - n].Next = nil
		resultListNode = head
	} else {
		addressMap[index - n].Next = addressMap[index - n + 2]
		resultListNode = head
	}

	return resultListNode

}