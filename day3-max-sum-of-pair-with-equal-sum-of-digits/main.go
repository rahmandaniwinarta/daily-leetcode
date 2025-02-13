package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 2, 3, 4, 4, 4, 5}
	nums1 := []int{27, 18, 43, 36, 13, 7}

	// fmt.Println(maximumSum(nums))
	fmt.Println(maximumSum(nums1))

}

func maximumSum(nums []int) int {
	res := -1

	sumMap := make(map[int]int)

	for _, num := range nums {
		fmt.Println(sumMap)
		sumIdx := sumIndex(num)
		if sumMap[sumIdx] > 0 {
			res = maxSum(res, sumMap[sumIdx]+num)
		}
		if sumMap[sumIdx] < num {
			sumMap[sumIdx] = num
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
