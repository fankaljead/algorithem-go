/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : 24_points.go
#   Last Modified : 2018-12-20 22:42
#   Describe      : 用户输入4个个位数（0-9），输出四则运算表达式，其结果为24
#
# ====================================================*/

package traverse

import (
	"log"
)

const (
	// PLUS meaning +
	PLUS = iota
	// SUBSTRACT meaning -
	SUBSTRACT
	// MULTIPLY meaning *
	MULTIPLY
	// DIVIDE meaning /
	DIVIDE
)

// CalculatePoints calculate a b c d into number points
func CalculatePoints(a, b, c, d, number int) {
	log.Println()
}
