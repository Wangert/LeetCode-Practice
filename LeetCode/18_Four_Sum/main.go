package main

import (
	"sort"
	"fmt"
	"strconv"
)

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}
	fmt.Println(fourSum(nums, 0))

}

func fourSum(nums []int, target int) [][]int {
	//对数组进行排序
	sort.Ints(nums)

	//定义结果数组
	results := make([][]int, 0)
	//计算去重后的数字个数
	length := len(nums)

	if length < 4 {
		return results
	}

	//用于去重的map
	deleteRepeatMap := make(map[string]string)

	//固定内侧两值
	for left1 := 1; left1 < length - 2; left1++ {

		for right1 := left1 + 1; right1 < length - 1; right1++ {
			//外侧左右位置
			left2, right2 := left1 - 1, right1 + 1

			for {

				if nums[left1] + nums[right1] + nums[left2] + nums[right2] > target {
					left2--
				} else if nums[left1] + nums[right1] + nums[left2] + nums[right2] < target {
					right2++
				} else {

					indexStr := strconv.Itoa(nums[left2]) + strconv.Itoa(nums[left1]) + strconv.Itoa(nums[right1]) + strconv.Itoa(nums[right2])
					temp := []int{nums[left2], nums[left1], nums[right1], nums[right2]}

					//去重
					if deleteRepeatMap[indexStr] != indexStr {

						deleteRepeatMap[indexStr] = indexStr
						results = append(results, temp)
						fmt.Println(strconv.Itoa(left2) + "|" + strconv.Itoa(left1) + "|" + strconv.Itoa(right1) + "|" + strconv.Itoa(right2) + "|")

					}

					for {

						left2--
						right2++

						if left2 < 0 || right2 >= length {
							break
						}

						if nums[left2] != temp[0] || nums[right2] != temp[3] {
							break
						}
					}
				}

				if left2 < 0 || right2 >= length {
					break
				}
			}
		}
	}

	return results

}
