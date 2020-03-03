package main

import "fmt"

/**
中心扩散法
 */

func main() {

	s := "babad"
	//s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(s))
}


func longestPalindrome(s string) string {

	if s == "" {
		return ""
	}

	//将字符串转换成rune切片
	runeStr := []rune(s)
	//获取字符串长度
	strLen := len(runeStr)
	//记录最大长度
	maxLen := 0
	//用于记录最长回文子串的左右位置
	var right, left int
	//游标
	i := 0


	for {

		if i >= strLen {
			break
		}

		//获得奇情况的左右位置和长度
		oddR, oddL, oddLen := longestByExpandAroundCenter(runeStr, i-1, i+1)
		evenR, evenL, evenLen := longestByExpandAroundCenter(runeStr, i, i+1)

		//判断当前最长
		if oddLen >= evenLen && oddLen > maxLen {

			maxLen = oddLen
			right = oddR
			left = oddL

		} else if oddLen < evenLen && evenLen > maxLen {

			maxLen = evenLen
			right = evenR
			left = evenL

		}

		i++

	}

	return string(runeStr[right+1:left])

}

//获取从中心扩散的最长回文子串
func longestByExpandAroundCenter(runeStr []rune, right, left int) (int, int, int) {
	//记录当前左右位置
	r, l := right, left

	for {
		//判断是否已经到边界
		if r < 0 || l >= len(runeStr) {
			break
		}

		//如果相等，向外扩散
		if runeStr[r] == runeStr[l] {
			r--
			l++
		} else {
			break
		}

	}

	return r, l, l - r - 1

}
