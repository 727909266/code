package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i :=1; i<len(nums); i ++ {
		if nums[index] != nums[i] {
			index ++
			nums[index] = nums[i]
		}
	}
	return index +1
}

func main()  {

}