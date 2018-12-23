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

import (
// "log"
)

// FindLongestPalindrome is to find the longest sub palindrome string in chars
// FindLongestPalindrome 找到 chars 中的最长回味 子序列
func FindLongestPalindrome(chars string) []string {
	result := make([]string, 0)    // 保存最长的回文子序列的string 数组
	resultMap := make(map[int]int) // 保存最长的回文子序列的下标对应 从start 到 end

	resultMap[0] = 0 // 初始化 最长回文子序列为字符串第一个元素
	for i := 0; i < len(chars); i++ {
		resultMap[i] = i
	}
	for index := 1; index < len(chars); index++ {
		for k, v := range resultMap {
			find(k, v, chars, resultMap, index)
		}
	}

	longestLength := findLongestLength(resultMap)
	result = findLongestSubString(longestLength, resultMap, chars)

	// log.Println(resultMap)
	// log.Println(result)

	return result
}

// find 找到从 start-(index-end) 到 index 是否为回文序列，是就将原来的删除，否则返回
func find(start, end int, chars string, resultMap map[int]int, index int) {
	if start < index-1 {
		for i := 0; i+start <= (start+index)/2; i++ {
			if chars[i+start] == chars[index-i] {
				continue
			} else {
				return
			}
		}
	}
	// delete(resultMap, start)
	resultMap[start] = index

}

func findLongestLength(resultMap map[int]int) int {
	longestLength := 1
	for k, v := range resultMap {
		temp := v - k
		if temp > longestLength {
			longestLength = temp
		}
	}

	return longestLength
}

func findLongestSubString(length int, resultMap map[int]int, chars string) []string {
	result := make([]string, 0)

	for k, v := range resultMap {
		temp := v - k
		if temp == length {
			result = append(result, chars[k:v+1])
		}

	}

	return result

}
