package main

/*
给定一个元素都是正整数的数组A ，正整数 L 以及 R (L <= R)。

求连续、非空且其中最大元素满足大于等于L 小于等于R的子数组个数。

例如 :
输入:
A = [2, 1, 4, 3]
L = 2
R = 3
输出: 3
解释: 满足条件的子数组: [2], [2, 1], [3].
注意:

L, R  和 A[i] 都是整数，范围在 [0, 10^9]。
数组 A 的长度范围在[1, 50000]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

import "fmt"

func numSubarrayBoundedMax(A []int, L int, R int) int {
	return itMaxCount(R, A) - itMaxCount(L - 1, A)
}


func itMaxCount(max int, a []int)(count int) {
	count = 0
	subCount := 0
	for _, v := range a {
		if  v <= max {
			subCount ++
			count += subCount
		} else {
			subCount = 0
		}
	}

	return

}

func main()  {
	fmt.Println(numSubarrayBoundedMax([]int{2,1,4,3}, 2, 3))

}