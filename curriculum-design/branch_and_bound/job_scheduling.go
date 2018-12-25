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

package branchAndBound

import (
	"log"
)

// NKScheduling 找到最佳调度
func NKScheduling(n, k int, t []int) int {
	len := new([]int)
	for i := 0; i < k; i++ {
		*len = append(*len, 0)
	}

	nkScheduling(0, len, &t, n, k)
	log.Println("len:", len)
	return *best

}

var (
	b    = 10000
	best *int
)

func init() {
	best = &b
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
func comp(len *[]int) int {
	tmp := 0
	for _, v := range *len {
		if v > tmp {
			tmp = v
		}
	}
	return tmp
}
