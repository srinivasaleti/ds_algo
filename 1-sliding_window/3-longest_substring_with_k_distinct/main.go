package main

import (
	"fmt"
)

// Problem Statement:
// Given a string, find the length of the longest substring in it
// with no more than K distinct characters.
// You can assume that K is less than or equal to the length of the given string.

func main() {
	runTest("araaci", 2, 4, "Test 1")         // "araa"
	runTest("araaci", 1, 2, "Test 2")         // "aa"
	runTest("cbbebi", 3, 5, "Test 3")         // "cbbeb" or "bbebi"
	runTest("aaaa", 1, 4, "Test 4")           // "aaaa"
	runTest("abcadcacacaca", 3, 11, "Test 5") // "cadcacacaca"
	runTest("", 2, 0, "Edge Test - Empty string")
	runTest("abc", 0, 0, "Edge Test - k = 0")

	runTest("", 2, 0, "Edge Test - Empty string")
	runTest("abc", 0, 0, "Edge Test - k = 0")
	runTest("abc", 10, 3, "Edge Test - k > string length")
	runTest("a", 1, 1, "Edge Test - single character string")
	runTest("aabacbebebe", 3, 7, "Test 6")  // "cbebebe"
	runTest("abaccc", 2, 4, "Test 7")       // "accc"
	runTest("abaccc", 3, 6, "Test 8")       // "abaccc"
	runTest("eceba", 2, 3, "Test 9")        // "ece"
	runTest("aaaaabbbbb", 2, 10, "Test 10") // "aaaaabbbbb"
	fmt.Println("✅ All tests passed!")
	fmt.Println("✅ All tests passed!")
}

func runTest(str string, k int, expected int, testName string) {
	result := longestSubstringKDistinct(str, k)
	if result != expected {
		panic(fmt.Sprintf("%s failed: expected %d, got %d", testName, expected, result))
	} else {
		fmt.Println(testName + " passed.")
	}
}

func longestSubstringKDistinct(str string, k int) int {
	frequency := map[string]int{}
	start := 0
	maxLen := 0
	for end := 0; end < len(str); end++ {
		chrEnd := string(str[end])
		frequency[chrEnd] = frequency[chrEnd] + 1

		for len(frequency) > k {
			chrStart := string(str[start])
			frequency[chrStart] = frequency[chrStart] - 1
			if frequency[chrStart] <= 0 {
				delete(frequency, chrStart)
			}
			start = start + 1
		}
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

// Algo: Longest Substring with K Distinct Characters
//
// 1. Use a map to count characters in the window.
// 2. Expand the window by moving 'end'.
// 3. If map has more than K chars, shrink window from 'start'.
// 4. Delete char from map if its count is zero.
// 5. Keep track of max window size.
