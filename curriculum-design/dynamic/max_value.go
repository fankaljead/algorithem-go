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

// FindMaxValue1 找到给定积分所能够获得的最大价值
// prices 商品价格列表, credits 商品所需积分列表, exchangeRates 商品积分价格兑换比率, totalCredit 总积分
// 返回商品最大价值时选中的商品列表下标
// 每种商品可以选择多次
func FindMaxValue1(prices, credits []float64, totalCredit float64) (float64, *[]float64, *[]float64) {

	weights := new([]float64) // 每个商品可以选择多次
	values := new([]float64)
	maxValue := float64(0)
	resultValues := new([]float64)
	resultWeights := new([]float64)

	for i := 0; i < len(prices); i++ {
		for j := 0; j <= int(totalCredit/credits[i]); j++ {
			*weights = append(*weights, credits[i])
			*values = append(*values, prices[i])
		}
	}

	log.Println("weights length:", len(*weights))
	log.Println("weights :", (*weights))
	f := new([][]float64)

	for i := 0; i < len(*values)+1; i++ {
		tf := new([]float64)
		for j := 0; j < int(totalCredit)+1; j++ {
			*tf = append(*tf, -1)
		}
		*f = append(*f, *tf)
	}

	result := new([]int)

	tracebackMemo(len(*values), int(totalCredit), weights, values, f, result)
	// tracebackMemo(len(prices), int(totalCredit), &credits, &prices, f, result)

	for _, v := range *result {
		maxValue += (*values)[v]
	}

	for _, v := range *result {
		*resultWeights = append(*resultWeights, (*weights)[v])
		*resultValues = append(*resultValues, (*values)[v])
	}

	return maxValue, resultWeights, resultValues
}

// FindMaxValue2 找到给定积分所能够获得的最大价值
// prices 商品价格列表, credits 商品所需积分列表, exchangeRates 商品积分价格兑换比率, totalCredit 总积分
// 返回商品最大价值时选中的商品列表下标
// 每种商品只能选择一次
func FindMaxValue2(prices, credits []float64, totalCredit float64) (float64, *[]int) {
	maxValue := float64(0)

	f := new([][]float64)

	for i := 0; i < len(prices)+1; i++ {
		tf := new([]float64)
		for j := 0; j < int(totalCredit)+1; j++ {
			*tf = append(*tf, -1)

		}
		*f = append(*f, *tf)

	}

	result := new([]int)

	// tracebackMemo(len(prices), int(totalCredit), weights, values, f)
	tracebackMemo(len(prices), int(totalCredit), &credits, &prices, f, result)

	for _, v := range *result {
		maxValue += prices[v]

	}

	return maxValue, result

}

func knapsackMemo(i, j int, weights, values *[]float64, f *[][]float64) float64 {
	if i == 0 || j == 0 {
		return 0
	}
	if (*f)[i][j] < 0 {
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
		if knapsackMemo(i, j, weights, values, f) ==
			(*values)[i-1]+knapsackMemo(i-1, j-int((*weights)[i-1]), weights, values, f) {
			*result = append(*result, i-1)
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
