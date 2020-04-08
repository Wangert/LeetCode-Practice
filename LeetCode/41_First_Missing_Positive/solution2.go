package main

import (
	"math"
	"fmt"
)

func main()  {

	nums := []int{3,4,-1,1}
	fmt.Println(firstMissingPositive2(nums))

}

func firstMissingPositive2(nums []int) int {

	if len(nums) == 0 {
		return 1
	}

	if len(nums) == 1 {
		 if nums[0] != 1 {
		 	return 1
		 } else {
		 	return 2
		 }
	}

	//判断是否有1
	flag := false
	for _, num := range nums {
		if num == 1 {
			flag = true
		}
	}

	if !flag {
		return 1
	}

	for i, num := range nums {

		if num <= 0 {
			nums[i] = 1
		}

		if num > len(nums) {
			nums[i] = 1
		}
	}

	for _, num := range nums {

		if int(math.Abs(float64(num))) == len(nums) {
			nums[0] *= -1
			continue
		}

		if nums[int(math.Abs(float64(num)))] > 0 {
			nums[int(math.Abs(float64(num)))] *= -1
		}

	}

	for i := 1; i < len(nums); i++ {

		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums) + 1

}