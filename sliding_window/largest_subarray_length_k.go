package main

// problems: https://github.com/Chanda-Abdul/Several-Coding-Patterns-for-Solving-Data-Structures-and-Algorithms-Problems-during-Interviews/blob/main/%E2%9C%85%20%20Pattern%2001%20%3A%20Sliding%20Window.md#maximum-sum-subarray-of-size-k-easy
func maxSubArrayOfSizeK(arr []int, k int) int {
	start := 0
	windowSum := 0
	maxSum := -9999999

	for end := 0; end < len(arr); end++ {
		windowSum += arr[end]

		if end >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[start]
			start++
		}
	}

	return maxSum
}
