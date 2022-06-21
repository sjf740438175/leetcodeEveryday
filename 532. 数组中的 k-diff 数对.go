package main

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	var count int
	for j:=n-1;j>=0;j--{
		if j < n-1 && nums[j] == nums[j+1]{
			continue
		}
		i := sort.SearchInts(nums,nums[j]-k)
		if nums[j]-nums[i] == k && i != j{
			count++
		}
	}
	return count
}
