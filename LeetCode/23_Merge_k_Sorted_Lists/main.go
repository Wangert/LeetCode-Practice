package main

import (
	"fmt"
	"math"
)

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */

//链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {

	test1 := &ListNode{1, nil}
	test1.Next = &ListNode{4, nil}
	test1.Next.Next = &ListNode{5, nil}

	test2 := &ListNode{1, nil}
	test2.Next = &ListNode{4, nil}
	test2.Next.Next = &ListNode{4, nil}

	test3 := &ListNode{2, nil}
	test3.Next = &ListNode{6, nil}

	lists := []*ListNode{test1, test2, test3}

	result := mergeKLists(lists)

	for {

		if result == nil {
			break
		}

		fmt.Println(result.Val)
		result = result.Next
	}


}


func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}


	//定义一个邻接表
	adList := make(map[int][]*ListNode)
	//记录最大值
	maxNum := math.MinInt64
	//记录最小值
	minNum := math.MaxInt64

	for _, list := range lists {

		for {

			if list == nil {
				break
			}

			if adList[list.Val] ==  nil {
				adList[list.Val] = make([]*ListNode, 2)
				adList[list.Val][0] = &ListNode{list.Val, nil}
				//记录该链表的尾部节点
				adList[list.Val][1] = adList[list.Val][0]
			} else {
				//将该数插入到对应链表的头部
				tempList := &ListNode{list.Val, adList[list.Val][0]}
				adList[list.Val][0] = tempList
			}

			if list.Val > maxNum {
				maxNum = list.Val
			}

			if list.Val < minNum {
				minNum = list.Val
			}

			list = list.Next

		}
	}

	if maxNum == math.MinInt64 && minNum == math.MaxInt64 {
		return nil
	}

	//记录结果链表头节点
	result := adList[minNum][0]

	//定义记录前一个数字
	preNum := minNum

	for i := minNum + 1; i <= maxNum; i++ {

		if adList[i] == nil {
			continue
		}

		adList[preNum][1].Next = adList[i][0]

		preNum = i

	}

	return result

}