package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l := len(nums)
	start := 0
	end := l
	i := l / 2
	little := true
	for ; i >= start && i < end; i = (start + end) / 2 {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			end = i
			little = true
		}
		if nums[i] < target {
			start = i + 1
			little = false
		}
		if start >= end {
			break
		}
	}
	if little {
		return i
	}
	return i + 1
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
