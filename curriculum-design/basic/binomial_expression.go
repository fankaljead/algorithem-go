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
	"strconv"
	"time"
)

// CalculateBinomialRecursion is to calculate the binomial express parse n and k with recursion method
// 计算二项式公式 使用递归
func CalculateBinomialRecursion(k, n int) int {
	if k > 0 && n > 0 {
		if k == 1 && n == 1 {
			return 1
		} else if k == 1 {
			return n
		} else {
			return CalculateBinomialRecursion(k-1, n-1) + CalculateBinomialRecursion(k, n-1)
		}
	}
	return 0
}

// CalculateBinomialMemo is to calculate the binomial express parse n and k with memory method
// 计算二项式公式 使用备忘录方法
func CalculateBinomialMemo(k, n int) int {
	if k > 0 && n > 0 {

		if k == 1 && n == 1 {
			return 1
		} else if k == 1 {
			return n
		} else if cknv[strconv.Itoa(k)+","+strconv.Itoa(n)] == 0 {
			cknv[strconv.Itoa(k)+","+strconv.Itoa(n)] = CalculateBinomialRecursion(k-1, n-1) + CalculateBinomialRecursion(k, n-1)
		}
		return cknv[strconv.Itoa(k)+","+strconv.Itoa(n)]
	}
	return 0
}

// CalculateBinomialIteration is to calculate the binomial express parse n and k with iteration method
// 计算二项式公式 使用迭代方法
func CalculateBinomialIteration(k, n int) int {
	var kns [][]int
	for i := 1; i <= n; i++ {
		var s []int
		for j := 1; j <= k; j++ {
			if i < j {
				s = append(s, 0)
			} else if j == 1 {
				s = append(s, i)
			} else if i == j {
				s = append(s, 1)
			} else if i-j == 1 {
				s = append(s, i)
			} else {
				s = append(s, kns[i-2][j-2]+kns[i-2][j-1])
			}
		}
		kns = append(kns, s)
	}

	return kns[n-1][k-1]
}

// CalculateBinomialBrute to brute force to calculate binomial express k, n
// 计算二项式公式 使用数学方法
func CalculateBinomialBrute(k, n int) int {
	denomenator := 1 // 分母
	member := 1      // 分子

	for i := 0; i < k; i++ {
		denomenator *= (n - i)
		member *= (i + 1)
	}

	return denomenator / member
}

// GetKN return k, n, k < n <20
// 随机生成 k, v, k < n < 20
func GetKN() (k, n int) {
	rand.Seed(time.Now().Unix())
	k = rand.Intn(15) + 1
	n = rand.Intn(20) + 1
	if n/2 < k {
		k = n - k
	}
	// k = rand.Int() % 10
	// n = rand.Int() % 15
	if k > n {
		k, n = n, k
	}

	return
}

var cknv map[string]int

func init() {
	cknv = make(map[string]int)
}
