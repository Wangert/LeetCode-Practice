package main

import "fmt"

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{1,3}
	fmt.Println(searchInsert(nums, 2))

}

func searchInsert(nums []int, target int) int {

	//计算数组长度
	numsLen := len(nums)

	if numsLen == 0 {
		return 0
	}

	if target <= nums[0] {
		return 0
	}

	if target > nums[numsLen - 1] {
		return numsLen
	}

	//定义当前搜索的左右两端索引
	left, right := 0, numsLen - 1

	for {

		minIndex := (right - left)>>1 + left
		fmt.Println(minIndex)

		if nums[minIndex] < target {
			left = minIndex + 1

			if left > right {
				return left
			}

		} else if nums[minIndex] > target {
			right = minIndex - 1

			if left > right {
				return minIndex
			}
		} else {
			return minIndex
		}
	}
}