/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : 24_points_test.go
#   Last Modified : 2018-12-21 14:06
#   Describe      : 测试24点
#
# ====================================================*/

package traverse

import (
	"log"
	"testing"
)

// TestCalculatePoints test CalculatePoints
func TestCalculatePoints(t *testing.T) {
	log.Println("4,6,7,5 -> 24")
	CalculatePoints(4, 6, 7, 5, 24)

}
