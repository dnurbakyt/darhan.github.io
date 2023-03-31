package main

import "fmt"

func numProduct(nums []int) int {
	var product int = 1
	for i := 0; i < len(nums); i++ {
		product = product * nums[i]
	}
	return product
}

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		numbers[i] = a
	}
	fmt.Print(numProduct(numbers))
}
