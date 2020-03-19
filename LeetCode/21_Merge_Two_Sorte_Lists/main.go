package main

import "fmt"

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */

func main()  {

	test1 := &ListNode{1, nil}
	test1.Next = &ListNode{2, nil}
	test1.Next.Next = &ListNode{4, nil}

	test2 := &ListNode{1, nil}
	test2.Next = &ListNode{3, nil}
	test2.Next.Next = &ListNode{4, nil}

	result := mergeTwoLists(test1, test2)

	for {

		if result == nil {
			break
		}

		fmt.Println(result.Val)
		result = result.Next
	}

}


//链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {


	if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}

	//用于遍历
	templ1 := l1
	templ2 := l2

	//定义指向结果链表的指针
	result := &ListNode{}
	//用于结果链表循环
	p := result

	for {

		if templ1.Val <= templ2.Val {
			tempNode := &ListNode{templ1.Val, nil}
			p.Next = tempNode
			templ1 = templ1.Next
		} else {
			tempNode := &ListNode{templ2.Val, nil}
			p.Next = tempNode
			templ2 = templ2.Next
		}

		p = p.Next

		if templ1 == nil {
			p.Next = templ2
			break
		} else if templ2 == nil {
			p.Next = templ1
			break
		}
	}

	return result.Next

}