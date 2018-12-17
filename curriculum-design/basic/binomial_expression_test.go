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
)

// CalculateBinomialRecursionTest 测试二项式公式的计算 使用递归
func TestCalculateBinomialRecursion(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialRecursion(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialRecursion=%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}

// TestCalculateBinomialMemo测试二项式公式的计算 使用备忘录方法
func TestCalculateBinomialMemo(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialMemo(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialMemo=%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}

// TestCalculateBinomialIteration 测试二项式公式的计算 使用备忘录方法
func TestCalculateBinomialIteration(t *testing.T) {
	k, n := GetKN()
	brute := CalculateBinomialBrute(k, n)
	recursion := CalculateBinomialIteration(k, n)
	log.Printf("C(%d, %d)=CalculateBinomialBrute=%d, CalculateBinomialIteration =%d\n", k, n, brute, recursion)
	if brute != recursion {
		t.Error()
	}
}
