package main

import (
	"fmt"
	"strconv"
)

/**
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

//链表节点
type ListNode struct {
	//值
	Val int
	//下一个节点
	Next *ListNode

}

//计算两数之和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//结果指针
	result := &ListNode{Val:-1}
	//临时节点
	var tmpNode *ListNode
	//进位标识
	flag := 0
	//存储临时数值
	var tmpNum int

	//记录头节点，用于临时存储节点
	tmpNode = result

	for {
		//判断两数是否遍历结束
		if l1 == nil && l2 == nil {
			//判断是否有进位，若有则新增结果节点
			if flag == 1 {
				tmpNode.Next = &ListNode{Val:1}
			}

			result = result.Next

			break
		}

		//判断l1是否遍历结束，若l2未结束则将l1的节点值置为0
		if l1 == nil {
			l1 = &ListNode{Val:0}
		}
		//判断l2是否遍历结束，若l1未结束则将l1的节点值置为0
		if l2 == nil {
			l2 = &ListNode{Val:0}
		}

		//判断是否进位
		if l1.Val + l2.Val + flag >= 10 {
			//计算相加的未进位值
			tmpNum = l1.Val + l2.Val + flag - 10
			//进位
			flag = 1

		} else {
			//计算相加的未进位值
			tmpNum = l1.Val + l2.Val + flag
			//不进位
			flag = 0

		}
		//创建新结果节点
		tmpNode.Next = &ListNode{Val:tmpNum}
		tmpNode = tmpNode.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	return result

}

func main() {

	//准备两个数
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{8, nil}
	l2.Next = &ListNode{2, nil}
	l2.Next.Next = &ListNode{7, nil}

	result := addTwoNumbers(l1, l2)

	fmt.Println("Result:")
	fmt.Print("[")

	for {

		if result.Next == nil {
			fmt.Print(strconv.Itoa(result.Val) + "]")
			break
		}

		fmt.Print(strconv.Itoa(result.Val) + " ")

		result = result.Next
	}

}
