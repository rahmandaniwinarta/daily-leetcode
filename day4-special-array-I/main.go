package main

import "fmt"

func main() {
	fmt.Println(isArraySpecial([]int{2, 1, 4}))
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))
}

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == 0 && nums[i+1]%2 == 0 {
			return false
		}
		if nums[i]%2 != 0 && nums[i+1]%2 != 0 {
			return false
		}
	}
	return true
}
