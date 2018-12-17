/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : binomial_expression.go
#   Last Modified : 2018-12-17 11:35
#   Describe      : 计算二项式公式 基本的递归算法第 1 题
#
# ====================================================*/

package basic

import (
	"math/rand"
	"time"
)

// CalculateBinomialRecursion is to calculate the binomial express parse n and k with recursion method
func CalculateBinomialRecursion(k, n int) int {
	if k == 1 && n == 1 {
		return 1
	} else if k == 1 {
		return n
	} else {
		return CalculateBinomialRecursion(k-1, n-1) + CalculateBinomialRecursion(k, n-1)
	}
}

// CalculateBinomialMemo is to calculate the binomial express parse n and k with memory method
func CalculateBinomialMemo(k, n int) int {
	if k == 1 && n == 1 {
		return 1
	} else if k == 1 {
		return n
	} else {
		return CalculateBinomialRecursion(k-1, n-1) + CalculateBinomialRecursion(k, n-1)
	}
}

// CalculateBinomialBrute to brute force to calculate binomial express k, n
func CalculateBinomialBrute(k, n int) int {
	denomenator := 1
	member := 1

	for i := 0; i < k; i++ {
		denomenator *= (n - i)
		member *= (i + 1)
	}

	return denomenator / member
}

// GetKN return k, n, k < n <20
func GetKN() (k, n int) {
	rand.Seed(time.Now().Unix())
	k = rand.Intn(15) + 1
	n = rand.Intn(20) + 1
	// k = rand.Int() % 10
	// n = rand.Int() % 15
	if k > n {
		k, n = n, k
	}

	return
}
