package main

import "fmt"

/**
中心扩散法
 */

func main() {

	//s := "babad"
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
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
	var left, right int
	//游标
	i := 0


	for {

		if i >= strLen {
			break
		}

		//获得奇情况的左右位置和长度
		oddL, oddR, oddLen := longestByExpandAroundCenter(runeStr, i-1, i+1)
		evenL, evenR, evenLen := longestByExpandAroundCenter(runeStr, i, i+1)

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

	return string(runeStr[left+1:right])

}

//获取从中心扩散的最长回文子串
func longestByExpandAroundCenter(runeStr []rune, left, right int) (int, int, int) {
	//记录当前左右位置
	l, r := left, right

	for {
		//判断是否已经到边界
		if l < 0 || r >= len(runeStr) {
			break
		}

		//如果相等，向外扩散
		if runeStr[l] == runeStr[r] {
			l--
			r++
		} else {
			break
		}

	}

	return l, r, r - l - 1

}
