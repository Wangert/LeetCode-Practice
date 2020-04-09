package main

import "fmt"

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	height := []int{2,0,2}
	fmt.Println(trap(height))

}

func trap(height []int) int {
	
	if len(height) <= 2 {
		return 0
	}

	return nextLayerRainWater(height)
}

//计算下一层的雨水量
func nextLayerRainWater(updatedHeight []int) int {

	//记录当前雨水量
	currentRainWater := 0
	//定义左右游标
	left, right := 0, 1
	//统计非0的个数
	notZero := 0

	for {
		if updatedHeight[left] == 0 {
			left++
			right++
		} else {
			break
		}

		if left >= len(updatedHeight) - 1 {
			return 0
		}
	}

	updatedHeight[left]--
	notZero++

	//临时水量
	tempRainWater := 0

	for {

		if updatedHeight[right] != 0 {
			currentRainWater = currentRainWater + tempRainWater
			tempRainWater = 0
			updatedHeight[right]--
			left = right
			notZero++
		} else {
			tempRainWater++
		}

		right++

		if right > len(updatedHeight) - 1 {
			break
		}
	}

	if notZero <= 1 {
		return 0
	}

	return currentRainWater + nextLayerRainWater(updatedHeight)
}