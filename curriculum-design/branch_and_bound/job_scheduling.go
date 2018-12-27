/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : job_scheduling.go
#   Last Modified : 2018-12-24 14:15
#   Describe      : 分支限界，n 个任务 k 个可并行的工作的机器 第 2 题
#
# ====================================================*/

package branchandbound

import (
// "log"
)

// NKScheduling 找到最佳调度
func NKScheduling(n, k int, t []int) (int, [][]int) {
	len := new([]int)
	for i := 0; i < k; i++ {
		*len = append(*len, 0)
	}

	result := make(map[int]int)

	// nkScheduling(0, len, &t, n, k)
	nkScheduling2(0, len, &t, n, k, &result)

	lastResult := new([][]int)

	for i := 0; i < k; i++ {
		temp := new([]int)
		*lastResult = append(*lastResult, *temp)
	}

	for k := range finalMap {
		(*lastResult)[finalMap[k]] = append((*lastResult)[finalMap[k]], t[k])

	}
	return *best, *lastResult

}

var (
	b        = 10000
	best     *int
	finalMap map[int]int
)

func init() {
	best = &b
	finalMap = make(map[int]int)
}
func nkScheduling(dep int, len, t *[]int, n, k int) {
	if dep == n {
		tmp := comp(len)
		if tmp < *best {
			*best = tmp
		}
		return
	}
	for i := 0; i < k; i++ {
		(*len)[i] += (*t)[dep]
		if (*len)[i] < *best {
			nkScheduling(dep+1, len, t, n, k)
		}
		(*len)[i] -= (*t)[dep]
	}

}
func nkScheduling2(dep int, lens, t *[]int, n, k int, result *map[int]int) {
	if dep == n {
		tmp := comp(lens)
		if tmp < *best {
			*best = tmp
			for k, v := range *result {
				finalMap[k] = v
			}
		}
		return
	}
	for i := 0; i < k; i++ {
		(*lens)[i] += (*t)[dep]
		(*result)[dep] = i
		if (*lens)[i] < *best {
			nkScheduling2(dep+1, lens, t, n, k, result)
		}
		(*lens)[i] -= (*t)[dep]
		delete(*result, dep)
	}

}
func sum(a *[]int) int {
	sum := 0
	for _, v := range *a {
		sum += v
	}
	return sum
}
func comp(len *[]int) int {
	tmp := 0
	for _, v := range *len {
		if v > tmp {
			tmp = v
		}
	}
	return tmp
}
