/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : palindrome.go
#   Last Modified : 2018-12-22 14:20
#   Describe      : 回文 求出最长的回文子序列 动态规划
#
# ====================================================*/

package dynamic

// FindLongestPalindrome is to find the longest sub palindrome string in chars
// FindLongestPalindrome 找到 chars 中的最长回味 子序列
func FindLongestPalindrome(chars string) []string {
	result := make([]string, 0)
	resultMap := make(map[int]int)
	resultMap[0] = 0

	return result
}
