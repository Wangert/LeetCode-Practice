package main

import "fmt"

/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

 

示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

 

说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
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
	test1.Next.Next.Next.Next = &ListNode{5, nil}

	result := reverseKGroup(test1, 3)

	for {

		if result == nil {
			break
		}

		fmt.Println(result.Val)
		result = result.Next
	}

}


func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}

	//存储当前需反转节点
	currentReverNodes := make([]*ListNode, 0)
	//用于遍历链表
	tempList := head

	//初始化第一轮反转节点
	for i := 0; i < k; i++ {

		//如果第一轮不满k个节点，即返回原链表
		if tempList == nil {
			return head
		}

		currentReverNodes = append(currentReverNodes, tempList)
		tempList = tempList.Next

	}

	//定义前一组交换后的最后节点
	preNewLastNode := &ListNode{}

	//结果头指针
	result := currentReverNodes[k - 1]

	for {
		//当前首节点指向当前最后节点所指向的节点
		currentReverNodes[0].Next = currentReverNodes[k - 1].Next

		//其他节点向前指向
		for i := k - 1; i > 0; i-- {
			currentReverNodes[i].Next = currentReverNodes[i - 1]
		}

		//将前一组反转后的最后节点指向当前组反转后的首节点
		preNewLastNode.Next = currentReverNodes[k - 1]
		//改变前一组反转后的最后节点，供下次循环使用
		preNewLastNode = currentReverNodes[0]


		tempList = currentReverNodes[0].Next
		//初始化第一轮反转节点
		for i := 0; i < k; i++ {

			if tempList == nil {
				goto END
			}

			currentReverNodes[i] = tempList
			tempList = tempList.Next

		}

	}

END:
	return result

}
