// Problem:
// Given a **sorted array** of integers, find the **indices of the two numbers**
// such that they add up to a given **target** number.
// Return the indices as a slice of 2 integers.
// If no such pair exists, return nil.
//
// Example:
// Input: arr = [1, 2, 3, 4, 6], target = 6
// Output: [1, 3]  // Because 2 + 4 = 6
//
// Constraints:
// - Input array is sorted in ascending order
// - Return indices of the numbers (0-based)
// - Use the two-pointer approach for efficient solution

package main

import (
	"fmt"
)

func main() {
	runTest([]int{1, 2, 3, 4, 6}, 6, []int{1, 3}, "Test 1")
	runTest([]int{2, 5, 9, 11}, 11, []int{0, 2}, "Test 2")
	runTest([]int{1, 3, 5, 7}, 10, []int{1, 3}, "Test 3")
	runTest([]int{1, 2, 3, 9}, 8, nil, "Test 4 - No pair")
	runTest([]int{}, 5, nil, "Test 5 - Empty array")
	runTest([]int{1}, 1, nil, "Test 6 - Single element")
	fmt.Println("✅ All tests passed!")
}

func runTest(arr []int, target int, expected []int, name string) {
	result := pairWithTargetSum(arr, target)
	if !equal(result, expected) {
		panic(fmt.Sprintf("%s failed: expected %v, got %v", name, expected, result))
	} else {
		fmt.Println(name + " passed.")
	}
}

func equal(a, b []int) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TODO: Fill this function using two pointers logic
func pairWithTargetSum(arr []int, target int) []int {
	result := []int{}
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == target {
			result = append(result, i)
			result = append(result, j)
			return result
		}
		if arr[i]+arr[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

// Very short algo:
// 1. Set two pointers: i = 0, j = end of array
// 2. Loop while i < j:
//    - if arr[i] + arr[j] == target → return [i, j]
//    - if sum > target → move j left
//    - if sum < target → move i right
// 3. If no pair found, return nil
