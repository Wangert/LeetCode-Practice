package main

import "fmt"

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums := []int{2, 3, 1}
	nextPermutation(nums)

	fmt.Println(nums)

}

func nextPermutation(nums []int) {

	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	//标记被替换位置
	replaceIndex := -1
	//遍历游标,从最后一个数字开始
	i := len(nums) - 1

	for {

		if i - 1 < 0 {
			break
		}

		//当前是从大到小排列
		if nums[i] <= nums[i - 1] {
			i--
			continue
		}

		//此处是需被替换位置
		replaceIndex = i - 1
		break
	}


	if replaceIndex == -1 {
		reverseIntArray(nums)
	} else {

		for j := replaceIndex + 1; j < len(nums); j++ {

			if nums[replaceIndex] >= nums[j] {
				nums[replaceIndex], nums[j - 1] = nums[j - 1], nums[replaceIndex]
				reverseIntArray(nums[replaceIndex + 1:])
				break
			}

			if j == len(nums) - 1 {
				nums[replaceIndex], nums[j] = nums[j], nums[replaceIndex]
				reverseIntArray(nums[replaceIndex + 1:])
				break
			}
		}
	}
}

// 反转数组
func reverseIntArray(nums []int) []int {

	for from, to := 0, len(nums)-1; from < to; from, to = from + 1, to - 1 {
		nums[from], nums[to] = nums[to], nums[from]
	}
	return nums
}