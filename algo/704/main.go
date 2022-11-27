package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	i := len(nums) / 2
	if nums[i] > target {
		sub_r := search(nums[0:i], target)
		return sub_r
	}
	if nums[i] < target {
		sub_r := search(nums[i:len(nums)], target)
		if sub_r != -1 {
			return sub_r + i
		}
		return -1
	}
	if nums[i] == target {
		return i
	}
	return -1
}

func searchWhile(nums []int, target int) int {
	l := len(nums)
	start := 0
	end := l
	for i := l / 2; i >= start && i < end; i = (start + end) / 2 {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			end = i
		}
		if nums[i] < target {
			start = i + 1
		}
		if start >= end {
			break
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 5}
	target := 5
	fmt.Println(search(nums, target))
}
