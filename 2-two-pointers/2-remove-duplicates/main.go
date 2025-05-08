package main

import (
	"fmt"
)

// Problem:
// Given a sorted array of integers,
// remove duplicates **in-place** and return the count of unique elements.
//
// Example:
// Input:  [2, 3, 3, 3, 6, 9, 9]
// Output: 4 (unique elements: 2, 3, 6, 9)

func main() {
	runTest([]int{2, 3, 3, 3, 6, 9, 9}, 4, "Test 1")
	runTest([]int{2, 2, 2, 11}, 2, "Test 2")
	runTest([]int{1, 1, 1, 1, 1}, 1, "Test 3")
	runTest([]int{1, 2, 3, 4}, 4, "Test 4")
	runTest([]int{}, 0, "Test 5 - Empty array")
	fmt.Println("✅ All tests passed!")
}

func runTest(nums []int, expected int, testName string) {
	count := removeDuplicates(nums)
	if count != expected {
		panic(fmt.Sprintf("%s failed: expected %d, got %d", testName, expected, count))
	}
	fmt.Printf("%s passed. Unique count = %d\n", testName, count)
}

// Function to remove duplicates in-place and return count of unique elements
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := 0
	for j < len(nums) {
		if nums[i] != nums[j] {
			i = i + 1
			nums[i] = nums[j]
		}
		j = j + 1
	}

	return i + 1
}
