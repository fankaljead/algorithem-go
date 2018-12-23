/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : palindrome_test.go
#   Last Modified : 2018-12-22 20:17
#   Describe      : 动态规划 找出最长子序列 测试
#
# ====================================================*/

package dynamic

import (
	"log"
	"testing"
)

// TestFindLongestPalindrome test the FindLongestPalindrome function
// 测试 FindLongestPalindrome 函数 找到最长回文子序列
func TestFindLongestPalindrome(t *testing.T) {

	log.Println("最长回文子序列")

	test1 := "ACA"
	result1 := FindLongestPalindrome(test1)
	log.Println("原字符串为:", test1)
	log.Println("最长回文子序列为:", result1)
	log.Println()

	test2 := "ACACCAB"
	result2 := FindLongestPalindrome(test2)
	log.Println("原字符串为:", test2)
	log.Println("最长回文子序列为:", result2)
	log.Println()

	test3 := "ACACCAB"
	result3 := FindLongestPalindrome(test3)
	log.Println("原字符串为:", test3)
	log.Println("最长回文子序列为:", result3)
	log.Println()

	test4 := "ACAACADBBACCAB"
	result4 := FindLongestPalindrome(test4)
	log.Println("原字符串为:", test4)
	log.Println("最长回文子序列为:", result4)

}
