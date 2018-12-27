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

import (
	"log"
)

// FindSupermarketLocation 找到超市位置
// matrix 各个点到其它点的距离 pNumber 每个点到超市的人次
func FindSupermarketLocation(matrix [][]int, pNumber []int) (position, length int, path [][]int) {

	// 点的下标 -> 路径
	//resultMap := make(map[int][][]int)

	// 遍历每个点
	for i := 0; i < len(matrix); i++ {
		path := new([][]int)
		selfToSelf := new([]int)
		*selfToSelf = append(*selfToSelf, i) // 自己到自己
		*path = append(*path, *selfToSelf)

		for j := 0; j < len(matrix); j++ {
			pathToJ := findShortestPath(i, j, matrix)
			*path = append(*path, *pathToJ)
		}
	}

	matrixS, resultMap := findShortest(matrix)

	log.Println(matrixS)
	log.Println(resultMap)
	return
}

// 找到最短路径
func findShortest(matrix [][]int) (matrixS *[][]int, pPath *map[int][][]int) {

	matrixS = new([][]int)
	path := make(map[int][][]int)
	// 复制原数组
	for i := 0; i < len(matrix); i++ {
		tP := new([]int)
		for j := 0; j < len(matrix[i]); j++ {
			*tP = append(*tP, matrix[i][j])
		}
		*matrixS = append(*matrixS, *tP)
	}

	// 找出最短路径
	for k := 0; k < len(*matrixS); k++ {
		kA := new([][]int)
		for i := 0; i < len(*matrixS); i++ {
			kTi := new([]int)
			for j := 0; j < len(*matrixS); j++ {
				minP, po := min((*matrixS)[i][j], (*matrixS)[i][k]+(*matrixS)[k][j])
				(*matrixS)[i][j] = minP
				if po == 1 {
					*kTi = append(*kTi, j)
				} else {
					*kTi = append(*kTi, k)
				}
			}
			*kA = append(*kA, *kTi)
		}
		path[k] = *kA
	}
	pPath = &path

	return matrixS, pPath
}

func min(a, b int) (int, int) {
	if a > b {
		return b, 2
	}
	return a, 1
}

func findShortestPath(i, j int, matrix [][]int) *[]int {
	path := new([]int)
	return path
}
