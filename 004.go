package main

import (
	"fmt"
	"math"
)

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func solve(nums1 []int, nums2 []int, k int) float64 {
	al :=0
	bl := 0
	m := len(nums1)
	n := len(nums2)
	for {
		if al >= m {
			fmt.Println(float64(nums2[bl + k]))
			return float64(nums2[bl + k])
		}
		if bl >= n {
			fmt.Println(float64(nums1[al + k]))
			return  float64(nums1[al + k])
		}
		amid := math.MaxInt64
		bmid := math.MaxInt64
		if al + k/2 < m {
			amid = nums1[al + k/2]
		}
		if bl + k/2 < n {
			bmid = nums2[bl + k/2]
		}
		if k == 0 {
			fmt.Println(math.Min(float64(amid), float64(bmid)))
			return math.Min(float64(amid), float64(bmid))

		}

		if amid < bmid {
			al = al + k/2 + 1
		} else {
			bl = bl + k/2 + 1
		}

		k = k - k/2
	}

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if (m + n) %2== 1 {
		return solve(nums1, nums2, (m + n)/2)
	} else {
		return (solve(nums1, nums2, (m + n)/2) + solve(nums1, nums2, (m + n)/2 -1))/2
	}



}
func main()  {
	a := []int{1,2}
	b := []int{3,4}
	fmt.Println(findMedianSortedArrays(a, b))
}
