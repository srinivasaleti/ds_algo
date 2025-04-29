package main

import (
	"fmt"
	"math"
)

/*
Problem:
---------
Given an array of **positive numbers** and a positive number **S**,
find the length of the smallest **contiguous subarray** whose sum is
greater than or equal to **S**.

If no such subarray exists, return **0**.

Examples:
---------
Input: nums = [2, 1, 5, 2, 3, 2], S = 7
Output: 2 → [5,2] or [2,3,2]

Input: nums = [2, 1, 5, 2, 8], S = 7
Output: 1 → [8]

Input: nums = [3, 4, 1, 1, 6], S = 8
Output: 3 → [1,1,6]

Edge Cases:
-----------
- If array is empty → return 0
- If all numbers are smaller than S and can't be added up to ≥ S → return 0
- If S is 0 → smallest subarray length is 1 (first element)
- If first number alone is ≥ S → return 1
- If S is very large → might return 0 if not possible
*/

func main() {
	runTest([]int{2, 1, 5, 2, 3, 2}, 7, 2, "Test 1")
	runTest([]int{2, 1, 5, 2, 8}, 7, 1, "Test 2")
	runTest([]int{3, 4, 1, 1, 6}, 8, 3, "Test 3")
	runTest([]int{1, 1, 1, 1, 1}, 5, 5, "Test 4")
	runTest([]int{1, 1, 1, 1, 1}, 6, 0, "Test 5 - No valid subarray")
	runTest([]int{10, 2, 3}, 6, 1, "Test 6 - Single large number")
	runTest([]int{}, 10, 0, "Test 7 - Empty array")
	runTest([]int{1, 2, 3, 4}, 0, 1, "Test 8 - S is zero")
	fmt.Println("✅ All tests passed!")
}

func runTest(nums []int, s int, expected int, testName string) {
	result := findMinSubArrayLen(s, nums)
	if result != expected {
		panic(fmt.Sprintf("%s failed: expected %d, got %d", testName, expected, result))
	} else {
		fmt.Println(testName + " passed.")
	}
}

func findMinSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if s == 0 {
		return 1
	}

	result := int(math.Inf(1))
	sum := 0
	i := 0
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		sum = sum + nums[windowEnd]
		for sum >= s {
			result = min(windowEnd-i+1, result)
			sum = sum - nums[i]
			i = i + 1
		}
	}
	if result == int(math.Inf(1)) {
		return 0
	}
	return result
}

// Algorithm to find smallest subarray length with sum >= S
//
// 1. If array is empty → return 0
// 2. If S is 0 → return 1 (any single number is enough)
//
// 3. Create two variables:
//    - sum → to store current window sum
//    - i   → to mark the start of the window
//    - result → store the smallest length (start with big number)
//
// 4. Loop from windowEnd = 0 to end of array:
//    - Add nums[windowEnd] to sum
//    - While sum is greater than or equal to S:
//       a. Update result with minimum of current length and result
//       b. Remove nums[i] from sum (shrink window from left)
//       c. Move i forward by 1
//
// 5. After loop, if result was not updated (still very big), return 0
// 6. Otherwise return result
