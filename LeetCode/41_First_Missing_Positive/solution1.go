package main

import (
	"sort"
	"fmt"
)

/**
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

 

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
 

提示：

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{3,4,-1,1}
	fmt.Println(firstMissingPositive(nums))

}

func firstMissingPositive(nums []int) int {
	//先从小到大排序
	sort.Ints(nums)
	//记录前一次的正数
	preNum := 0

	for _, num := range nums {

		if num <= 0 || preNum == num {
			continue
		}

		if preNum == 0 && num > 1 {
			return 1
		}

		if preNum + 1 != num {
			return preNum + 1
		} else {
			preNum = num
		}
	}

	return preNum + 1

}