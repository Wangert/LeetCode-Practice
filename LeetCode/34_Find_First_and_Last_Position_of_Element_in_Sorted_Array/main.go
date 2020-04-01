package main

import "fmt"

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func main()  {

	nums := []int{2,2}
	fmt.Println(searchRange(nums, 2))

}

func searchRange(nums []int, target int) []int {

	//数组长度
	numsLen := len(nums)

	if numsLen == 0 {
		return []int{-1, -1}
	}

	if target < nums[0] {
		return []int{-1, -1}
	}

	//存储结果
	result := []int{-1, -1}
	//定义搜索数组的左右索引
	left, right := 0, numsLen - 1

	for {

		midIndex := (right - left)>>1 + left
		//如果中间值等于目标值，往外扩散
		if nums[midIndex] == target {
			centerSpread(nums, result, midIndex, midIndex, target)
			goto END

		} else if nums[midIndex] > target {

			right = midIndex - 1

			if right < left {
				goto END
			}
		} else {

			left = midIndex + 1

			if left > right {
				goto END
			}
		}
	}


END:
	return result

}

func centerSpread(nums, result []int, i, j, target int) {

	numsLen := len(nums)

	for {
		if nums[i] == target {
			i--
			if i < 0 {
				result[0] = 0
				break
			}
		} else if nums[i] != target {
			result[0] = i + 1
			break
		}
	}

	for {
		if nums[j] == target {
			j++
			if j >= numsLen {
				result[1] = numsLen - 1
				break
			}
		} else if nums[j] != target {
			result[1] = j - 1
			break
		}
	}
}
