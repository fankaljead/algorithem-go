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

package branchAndBound

import (
	"log"
	"testing"
)

func TestNKScheduling(t *testing.T) {

	best := NKScheduling(7, 3, []int{2, 14, 4, 16, 6, 5, 3})
	log.Println("best:", best)
}
