package main

import (
	"fmt"
)

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start := 0
	end := 0
	for i:=0 ; i<len(s);i++ {
		len1 :=handle(s,i,i)
		len2 := handle(s,i,i+1)
		lene := len2
		if len1 > len2 {
			lene = len1
		}
		if lene > end - start {
			start = i - (lene-1)/2
			end = i + lene/2
		}
	}
	return s[start:end+1]

}


func handle(s string, l, r int) int {
	for {
		if l <0 ||r >=len(s)||s[l] != s[r]{
			break
		}
		l--
		r++
	}
	return r-l-1
}
func main()  {
	fmt.Println(longestPalindrome("aacdefcaa"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
