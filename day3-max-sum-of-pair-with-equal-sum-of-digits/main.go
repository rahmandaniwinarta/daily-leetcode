package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 4, 5}

	fmt.Println(maximumSum(nums))
}

func maximumSum(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if sumIndex(nums[i]) == sumIndex(nums[j]) {
				res = maxSum(res, nums[i]+nums[j])
			}
		}
	}
	return res

}

func sumIndex(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}

func maxSum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
