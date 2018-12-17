/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : binomial_expression_test.go
#   Last Modified : 2018-12-17 11:43
#   Describe      : 测试二项式公式的计算
#
# ====================================================*/

package basic

import (
	"log"
	"testing"
	"time"
)

// CalculateBinomialRecursionTest 测试二项式公式的计算 使用递归
func TestCalculateBinomialRecursion(t *testing.T) {
	for i := 10; i > 0; i-- {
		k, n := GetKN()
		brute := CalculateBinomialBrute(k, n)
		recursion := CalculateBinomialBrute(k, n)
		log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialRecursion=%d\n", k, n, brute, recursion)
		if brute != recursion {
			t.Error()
		}
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
