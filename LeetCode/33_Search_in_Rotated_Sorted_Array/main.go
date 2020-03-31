package main

import "fmt"

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(search(nums, 3))

}

func search(nums []int, target int) int {

	//数组长度
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}

	if target < nums[0] && target > nums[numsLen - 1] {
		return -1
	}

	if target < nums[0] {

		if nums[numsLen - 1] >= nums[0] {
			return -1
		}

		i := numsLen - 1

		for {

			if target == nums[i] {
				return i
			} else if target > nums[i] {
				return -1
			}

			if nums[i] < nums[i - 1] {
				return -1
			}

			i--
		}
	} else {

		i := 0

		for {

			if target == nums[i] {
				return i
			} else if target < nums[i] {
				return -1
			}

			if i == numsLen - 1 {
				return -1
			}
			if nums[i] > nums[i + 1] {
				return -1
			}

			i++
		}
	}
}