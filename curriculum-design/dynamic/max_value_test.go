/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : max_value_test.go
#   Last Modified : 2018-12-23 14:17
#   Describe      : 积分最大价值 测试
#
# ====================================================*/

package dynamic

import (
	"log"
	"testing"
)

// TestFindMaxValue test FindMaxValue
func TestFindMaxValue(t *testing.T) {
	prices := [...]float64{200, 400, 500, 250, 500, 60, 300, 200}          // 商品价格列表
	credits := [...]float64{5000, 1000, 2500, 500, 3000, 1000, 2000, 1200} // 商品所需积分列表
	totalCredit := float64(10000)                                          // 总积分

	maxValue, resultWeights, resultValues := FindMaxValue1(prices[:], credits[:], totalCredit)
	log.Println("考虑每种商品可以选择多次的情况")
	log.Println("总积分为:", totalCredit)
	log.Println("商品价格列表为: ", prices)
	log.Println("商品所需积分为: ", credits)
	log.Println("最大价值列表商品价格: ", *resultWeights)
	log.Println("最大价值列表商品积分: ", *resultValues)
	log.Println("最大价值为: ", maxValue)
	log.Println()

	maxValue, result := FindMaxValue2(prices[:], credits[:], totalCredit)
	log.Println("考虑每种商品只能选择一次的情况")
	log.Println("总积分为:", totalCredit)
	log.Println("商品价格列表为: ", prices)
	log.Println("商品积分列表为: ", credits)
	log.Println("最大价值列表商品选择的为: ", *result)
	log.Println("最大价值为: ", maxValue)
}
