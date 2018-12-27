/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : supermarket_location.go
#   Last Modified : 2018-12-27 14:11
#   Describe      : 超市选址 附加题目 3 带权有向图中心
#
# ====================================================*/

package extra

//FindSupermarketLocation 找到超市位置
func FindSupermarketLocation(matrix [][]int) (position, length int) {

	result := new([]int)

	for i := 0; i < len(matrix); i++ {
		sum := 0
		for j := 0; j < len(matrix[i]); j++ {
			if j != i {
				sum += matrix[i][j]
			}
		}
		*result = append(*result, sum)
	}

	maxIndex := 0
	for k, v := range *result {
		if v > (*result)[maxIndex] {
			maxIndex = k
		}
	}

	return maxIndex, (*result)[maxIndex]
}
