package main

import (
	"fmt"
)

/*
Problem Statement:
Given an integer array `nums` sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
Input: nums = [-4, -1, 0, 3, 10]
Output: [0, 1, 9, 16, 100]
Explanation: After squaring, the array becomes [16, 1, 0, 9, 100].
After sorting, it becomes [0, 1, 9, 16, 100].

Example 2:
Input: nums = [-7, -3, 2, 3, 11]
Output: [4, 9, 9, 49, 121]

Constraints:
- 1 <= nums.length <= 10^4
- -10^4 <= nums[i] <= 10^4
- `nums` is sorted in non-decreasing order.

*/

func main() {
	// Test cases
	runTest([]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}, "Test 1")
	runTest([]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}, "Test 2")
	runTest([]int{-2, -2, 0, 2, 2}, []int{0, 4, 4, 4, 4}, "Test 3")         // Case with multiple duplicates
	runTest([]int{-1, 0, 1}, []int{0, 1, 1}, "Test 4")                      // Case with small range of numbers
	runTest([]int{-10, -5, 0, 5, 10}, []int{0, 25, 25, 100, 100}, "Test 5") // Case with negative and positive values
	runTest([]int{}, []int{}, "Test 6")                                     // Edge case with an empty array
	runTest([]int{0, 0, 0}, []int{0, 0, 0}, "Test 7")                       // Case with all zeros
	runTest([]int{-100, 100}, []int{10000, 10000}, "Test 8")                // Large numbers
	runTest([]int{-1, 1, -1, 1}, []int{1, 1, 1, 1}, "Test 9")               // Case with alternating negative and positive values
	fmt.Println("✅ All tests passed!")
}

func runTest(nums []int, expected []int, testName string) {
	result := sortedSquares(nums)
	if !equal(result, expected) {
		panic(fmt.Sprintf("%s failed: expected %v, got %v", testName, expected, result))
	} else {
		fmt.Println(testName + " passed.")
	}
}

// Function to check if two arrays are equal
func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// TODO: Fill in the logic
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	j := firstPositiveIndex(nums)
	i := j - 1

	result := []int{}
	for i >= 0 && j < len(nums) {
		if (nums[i] * nums[i]) < (nums[j] * nums[j]) {
			result = append(result, nums[i]*nums[i])
			i = i - 1
		} else {
			result = append(result, nums[j]*nums[j])
			j = j + 1
		}
	}
	for i >= 0 {
		result = append(result, nums[i]*nums[i])
		i = i - 1
	}
	for j < len(nums) {
		result = append(result, nums[j]*nums[j])
		j = j + 1
	}

	// Your logic goes here
	return result
}

func firstPositiveIndex(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] < 0 {
			i = i + 1
		} else {
			return i
		}
	}
	return i
}

/**
This is a straightforward question. The only trick is that we can have negative numbers in the input array, which will make it a bit difficult to generate the output array with squares in sorted order.

An easier approach could be to first find the index of the first non-negative number in the array. After that, we can use Two Pointers to iterate the array. One pointer will move forward to iterate the non-negative numbers, and the other pointer will move backward to iterate the negative numbers. At any step, whichever number gives us a bigger square will be added to the output array.

Since the numbers at both ends can give us the largest square, an alternate approach could be to use two pointers starting at both ends of the input array. At any step, whichever pointer gives us the bigger square, we add it to the result array and move to the next/previous number according to the pointer.
**/
