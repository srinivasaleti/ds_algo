// Problem Statement:
//
// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value
// and return this value.
//
// Any answer with a calculation error less than 10^-5 will be accepted.
//
// Example:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75
// Explanation: Maximum average is from subarray [12,-5,-6,50]

package main

import (
	"fmt"
	"math"
)

func main() {
	runTest([]int{1, 12, -5, -6, 50, 3}, 10, 9.16, "Test 1")
	runTest([]int{1, 12, -5, -6, 50, 3}, 4, 12.75, "Test 2")
	runTest([]int{5, 5, 5, 5, 5}, 3, 5.0, "Test 3")
	runTest([]int{-10, -5, -3, -4, -2}, 2, -3.0, "Test 4")
	runTest([]int{1, 2, 3, 4, 5}, 5, 3.0, "Test 5")
	runTest([]int{2, 7, 1, 8, 2}, 1, 8.0, "Test 6")
	runTest([]int{}, 3, 0.0, "Edge Test 1: Empty array")
	runTest([]int{10, 20}, 0, 0.0, "Edge Test 2: k = 0")
	runTest([]int{10, 20}, -2, 0.0, "Edge Test 3: k < 0")
	runTest([]int{10, 20, 30}, 5, 20.0, "Edge Test 4: k > length of array")
	runTest([]int{0, 0, 0, 0, 0}, 3, 0.0, "Edge Test 5: All zero values")
	runTest([]int{-1000, 1000}, 1, 1000.0, "Edge Test 6: Max at end")
	runTest([]int{-1000, 1000}, 2, 0.0, "Edge Test 7: Mixed large range")
	fmt.Println("✅ All tests passed!")
}

func runTest(nums []int, k int, expected float64, testName string) {

	result := findMaxAverage(nums, k)
	if math.Abs(result-expected) > 0.01 {
		panic(fmt.Sprintf("%s failed: expected %.5f, got %.5f", testName, expected, result))
	} else {
		fmt.Println(testName + " passed.")
	}
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 || k <= 0 {
		return 0
	}
	if len(nums) <= k {
		return findSum(nums, len(nums)) / float64(len(nums))
	}
	sum := findSum(nums, k)
	maxAvg := sum / float64(k)
	for i := k; i < len(nums); i++ {
		sum = sum + float64(nums[i]) - float64(nums[i-k])

		newAvg := sum / float64(k)
		if newAvg > maxAvg {
			maxAvg = newAvg
		}
	}

	return maxAvg
}

func findSum(nums []int, k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		if i < len(nums) {
			sum = sum + float64(nums[i])
		}
	}
	return sum
}

/*
Algorithm Explanation:

1. **Initial Window Calculation**:
    - Start by calculating the sum of the first 'k' elements in the array. This will form the first window.

2. **Sliding Window**:
    - As we slide the window, we add the next element and subtract the first element of the previous window from the sum.
    - This allows us to compute the sum of the current window without iterating through all elements each time.

3. **Max Average Calculation**:
    - For each new window, we calculate the average by dividing the sum by 'k'.
    - We update the maximum average if the current window's average is higher.

4. **Edge Case Handling**:
    - If 'k' is less than or equal to zero, or if the array is empty, the function returns 0.
    - If 'k' is greater than or equal to the length of the array, the entire array is considered as one window.

Time Complexity:
- **O(n)**, where 'n' is the length of the array. We make only one pass through the array using the sliding window technique.

Example:
    Input: nums = [1, 12, -5, -6, 50, 3], k = 4
    Output: 12.75

    Explanation:
    - First window: [1, 12, -5, -6] → Average = 0.5
    - Second window: [12, -5, -6, 50] → Average = 12.75 (Max Average)
    - Third window: [-5, -6, 50, 3] → Average = 10.5

Thus, the maximum average is 12.75.
*/
