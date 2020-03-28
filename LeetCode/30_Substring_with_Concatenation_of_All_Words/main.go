package main

import "fmt"

/**
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

 

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	//s := "barfoothefoobarman"
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	//words := []string{"word","good","best","word"}

	//s1 := "wordgoodgoodgoodbestword"
	//words1 := []string{"word","good","best","good"}

	fmt.Println(findSubstring(s, words))

}

func findSubstring(s string, words []string) []int {

	//将字符串转换成rune切片
	sRunes := []rune(s)
	//字符串的长度
	sLength := len(s)
	//单词个数
	wLength := len(words)

	if wLength == 0 || sLength == 0 {
		return []int{}
	}

	//记录剩余单词的数量
	wordMaps := make(map[string]int)
	//用于循环判断
	tempWordMaps := make(map[string]int)

	//将单词组中单词出现的次数存在map中
	for _, word := range words {

		if wordMaps[word] == 0 {
			wordMaps[word] = 1
			tempWordMaps[word] = 1
			continue
		}

		wordMaps[word]++
		tempWordMaps[word]++
	}

	//单词长度
	wordLength := len(words[0])
	//定义窗口左右
	left, right := 0, wLength * wordLength
	//记录结果
	result := make([]int, 0)

	for {

		i := left
		j := right - wordLength

		for {

			//判断在剩余单词中是否存在
			if tempWordMaps[string(sRunes[i:i+wordLength])] != 0 {
				tempWordMaps[string(sRunes[i:i+wordLength])]--
				if i == j {
					goto IF
				} else if tempWordMaps[string(sRunes[j:j+wordLength])] != 0 {
					tempWordMaps[string(sRunes[j:j+wordLength])]--
				} else {
					left++
					right++
					goto NEXT
				}
			} else {
				left++
				right++
				goto NEXT
			}

		IF:
			if i + wordLength >= j {
				result = append(result, left)
				left++
				right++
				goto NEXT
			} else {
				i = i + wordLength
				j = j - wordLength
			}
		}

	NEXT:
		//复原
		for word, count := range wordMaps {
			tempWordMaps[word] = count
		}

		if right > sLength {
			break
		}
	}

	return result
}