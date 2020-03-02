package main

import (
	"fmt"
)

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
