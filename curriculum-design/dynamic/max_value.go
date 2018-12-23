/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : max_value.go
#   Last Modified : 2018-12-23 14:09
#   Describe      : 积分最大价值 动态规划第 6 题
#
# ====================================================*/

package dynamic

import (
	"log"
)

// FindMaxValue 找到给定积分所能够获得的最大价值
// prices 商品价格列表, credits 商品所需积分列表, exchangeRates 商品积分价格兑换比率, totalCredit 总积分
// 返回商品最大价值时选中的商品列表下标
func FindMaxValue(prices, credits, exchangeRates []float64, totalCredit float64) (float64, *[]int) {
	maxValue := float64(0)

	f := new([][]float64)

	for i := 0; i < len(prices)+1; i++ {
		tf := new([]float64)
		for j := 0; j < int(totalCredit)+1; j++ {
			*tf = append(*tf, -1)
		}
		*f = append(*f, *tf)
	}
	log.Println("f:", f)
	log.Println("f length:", len(*f))
	log.Println("f0 length:", len((*f)[0]))
	// weights := new([]int)
	// values := new([]int)
	result := new([]int)

	// tracebackMemo(len(prices), int(totalCredit), weights, values, f)
	tracebackMemo(len(prices), int(totalCredit), &credits, &prices, f, result)
	log.Println(prices)

	log.Println("f:", f)
	for _, v := range *result {
		maxValue += prices[v]
	}

	return maxValue, result
}

func knapsackMemo(i, j int, weights, values *[]float64, f *[][]float64) float64 {
	if i == 0 || j == 0 {
		return 0
	}
	if (*f)[i][j] == 0 {
		v := float64(0)
		if float64(j) < (*weights)[i-1] {
			v = knapsackMemo(i-1, j, weights, values, f)
		} else {
			v = max(knapsackMemo(i-1, j, weights, values, f),
				(*values)[i-1]+knapsackMemo(i-1, j-int((*weights)[i-1]), weights, values, f))
		}
		(*f)[i][j] = v
	}
	return (*f)[i][j]
}

func tracebackMemo(i, j int, weights, values *[]float64, f *[][]float64, result *[]int) {
	if i == 0 || j == 0 {
		return
	}
	if float64(j) >= (*weights)[i-1] {
		log.Println("(j-int((*weights)[i-1]) = ", (j - int((*weights)[i-1])))
		if knapsackMemo(i, j, weights, values, f) ==
			(*values)[i-1]+knapsackMemo(i-1, j-int((*weights)[i-1]), weights, values, f) {
			*result = append(*result, i-1)
			log.Println("i-1=", i-1)
			tracebackMemo(i-1, j-int((*weights)[i-1]), weights, values, f, result)
		} else {
			tracebackMemo(i-1, j, weights, values, f, result)
		}
	} else {
		tracebackMemo(i-1, j, weights, values, f, result)
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
