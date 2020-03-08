package main

import (
	"math"
	"fmt"
)

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。 

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {

	test := []int{2,3,10,5,7,8,9}
	fmt.Println(maxArea(test))


}

func maxArea(height []int) int {

	//两端游标
	left, right := 0, len(height) - 1
	//记录最大面积
	max := 0

	for {

		//计算left与right的面积
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))

		max = int(math.Max(float64(max), float64(area)))

		if height[left] < height[right] {
			left++
		}else {
			right--
		}

		if left >= right {
			break
		}

	}

	return max

}