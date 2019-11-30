package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
滑动窗口题目:

3. 无重复字符的最长子串

30. 串联所有单词的子串

76. 最小覆盖子串

159. 至多包含两个不同字符的最长子串

209. 长度最小的子数组

239. 滑动窗口最大值

567. 字符串的排列

632. 最小区间

727. 最小窗口子序列

作者：powcai
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */






func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}


func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]int, 0)
	res := 0
	left := 0
	for index, v := range s {
		_, ok := set[v]
		if ok {
			left = max(left, set[v] + 1)
		}
		set[v] = index
		res = max(res, index - left + 1)
	}
	return res
}
func main()  {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))

}
