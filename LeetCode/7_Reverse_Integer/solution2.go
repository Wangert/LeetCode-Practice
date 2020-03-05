package main

import "fmt"
import "math"

/**

方法二优于方法一

 */

func main()  {

	fmt.Println(reverse(-1234))

}

func reverse(x int) ( num int) {

	for x != 0 {
		num = num * 10 + x % 10
		x = x / 10
	}

	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

	return
}
