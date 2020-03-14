package main

import (
	"sort"
	"math"
	"fmt"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func main()  {

	//nums := []int{-1, 2, 1, -4}
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Println(threeSumClosest(nums, 82))

}

func threeSumClosest(nums []int, target int) int {

	//定义minSum
	minDistance := int(math.Abs(float64(nums[0] + nums[1] + nums[2] - target)))
	//记录结果和
	resultSum := nums[0] + nums[1] + nums[2]
	//从小到大排序
	sort.Ints(nums)
	//计算整数个数
	length := len(nums)

	for i := 1; i < length - 1; i++ {

		//定义当前起始左右位置
		left, right := i - 1, i + 1

		if nums[i] >= 0 && nums[i] + nums[0] - target > minDistance {
			continue
		}

		if nums[i] <= 0 && nums[i] + nums[length-1] - target < -minDistance {
			continue
		}


		for {

			if nums[i] + nums[left] + nums[right] - target >= minDistance {
				left--
			} else if nums[i] + nums[left] + nums[right] - target < -minDistance {
				right++
			} else {
				minDistance = int(math.Abs(float64(nums[i] + nums[left] + nums[right] - target)))

				resultSum = nums[i] + nums[left] + nums[right]

				if resultSum - target >=0 {
					left--
				} else {
					right++
				}

			}

			if left < 0 || right >= length {
				break
			}

		}

	}

	return resultSum

}