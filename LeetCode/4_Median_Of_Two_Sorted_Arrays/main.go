package main

import (
	"fmt"
)

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	nums1 := []int{1, 3, 5, 9}
	nums2 := []int{4, 7}

	fmt.Println(findMedianSortedArrays(nums1, nums2))


}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//获取总共有几个数字
	numberOfElements := len(nums1) + len(nums2)
	//计算两个数字的长度
	nums1Len, nums2Len := len(nums1), len(nums2)
	//中位数位置
	minIndex := numberOfElements / 2
	//定义遍历索引
	i, j := 0, 0
	//计数器
	count := 0
	//定义临时合并数组
	mergeArray := make([]int, 0)

	for {

		if i < nums1Len && (j >= nums2Len || nums1[i] <= nums2[j]) {
			mergeArray = append(mergeArray, nums1[i])
			i++
			count++
		}else {
			mergeArray = append(mergeArray, nums2[j])
			j++
			count++
		}

		//判断是否以遍历到中位数位置
		if count - 1 == minIndex {
			//判断奇偶
			if oddOrNot(numberOfElements) {
				return float64(mergeArray[minIndex])
			}else {
				return float64(mergeArray[minIndex-1] + mergeArray[minIndex]) / 2.0
			}
		}
	}
}

//判断奇偶,ture为奇数，false为偶数
func oddOrNot(num int) bool {
	return (num & 1) != 0
}
