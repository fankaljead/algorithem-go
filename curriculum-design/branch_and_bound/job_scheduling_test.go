/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : job_scheduling_test.go
#   Last Modified : 2018-12-24 14:18
#   Describe      :
#
# ====================================================*/

package branchandbound

import (
	"log"
	"testing"
)

func TestNKScheduling(t *testing.T) {
	best, result := NKScheduling(7, 3, []int{2, 14, 23, 11, 9, 3, 4, 16, 6, 5, 3})
	log.Println("最短作业时间为:", best)
	log.Println("分配情况如下:", result)
}
