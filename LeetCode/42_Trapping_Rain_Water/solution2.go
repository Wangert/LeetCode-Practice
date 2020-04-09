package main

import "fmt"

func main()  {

	height := []int{2,0,2}
	fmt.Println(trap2(height))

}

func trap2(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	//左右游标
	left, right := 0, len(height) - 1
	//分别记录双向的当前最高柱子
	leftMax, rightMax := 0, 0
	//分别记录双向覆盖面积
	leftArea, rightArea := 0, 0

	for {

		if height[left] > leftMax {
			leftMax = height[left]
		}

		leftArea = leftArea + leftMax


		if height[right] > rightMax {
			rightMax = height[right]
		}

		rightArea = rightArea + rightMax

		left++
		right--

		if right < 0 {
			break
		}
	}

	//柱子面积
	barArea := 0

	for _, h := range height {
		barArea = barArea + h
	}

	//计算矩形面积
	recArea := leftMax * len(height)

	return leftArea + rightArea - recArea - barArea

}
