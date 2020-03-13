package main

import (
	"fmt"
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{-1, 0, 0, 0, 1, 2, -1, -4}
	//nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))

}


func threeSum(nums []int) [][]int {
	//对数组进行排序
	sort.Ints(nums)

	//定义结果数组
	results := make([][]int, 0)
	//计算去重后的数字个数
	length := len(nums)

	if length < 3 {
		return results
	}

	//记录第几个0
	flag := 0

	//外层表示三者中间的值
	for i := 1; i < length - 1; i++ {
		//起始的左右位置
		left, right := i - 1, i + 1

		//以该数为中心，不可能有三数和为0
		if nums[i] + nums[length - 1] < 0 || nums[i] + nums[0] > 0 {
			continue
		}

		if nums[i] < 0 && nums[i] == nums[i + 1] || nums[i] > 0 && nums[i - 1] == nums[i]{
			continue
		}

		//考虑0为中心的情况
		//若第一次循环中心就是0，则需要判断左边是否也为0
		if nums[i] == 0 && flag == 0 && nums[left] == 0 {
			flag++
		}
		//数组中只有一个零
		if nums[i] == 0 && flag == 0 && nums[right] != 0 {
			flag++
			goto NEXT
		}

		//数组中有多个零，选第二个零
		if nums[i] == 0 {
			flag++

			if flag != 2 {
				continue
			}
		}
NEXT:
		for {

			if nums[left] + nums[i] + nums[right] > 0 {
				left--
			} else if nums[left] + nums[i] + nums[right] < 0 {
				right++
			} else {
				temp := []int{nums[left], nums[i], nums[right]}
				results = append(results, temp)

				for {

					left--
					right++

					if left < 0 || right >= length {
						break
					}

					if nums[left] != nums[left + 1] || nums[right] != nums[right - 1] {
						break
					}
				}
			}

			if left < 0 || right >= length {
				break
			}
		}
	}

	return results

}

/*
//数组去重
func DeleteRepeat(list []int) []int {
	//定义map
	mapdata := make(map[int]interface{})

	if len(list) <= 0 {
		return nil
	}

	//将原数组中的值作为map的键
	for _, v := range list {
		mapdata[v] = "true"
	}

	var datas []int

	for k, _ := range mapdata {

		datas = append(datas, k)
	}

	return datas
}
*/