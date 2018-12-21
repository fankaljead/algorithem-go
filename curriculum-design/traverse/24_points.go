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
	"strconv"
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

	numbers := [...]int{a, b, c, d}
	// 产生数的排列
	resultList := new(NumberList)
	Perm(numbers[:], 0, len(numbers)-1, resultList)
	// log.Println("resultList:", resultList)

	// 产生运算符的排列
	operators := [...]int{PLUS, SUBSTRACT, MULTIPLY, DIVIDE}
	operatorList := new(NumberList)
	Perm(operators[:], 0, len(operators)-1, operatorList)
	// log.Println("operatorList:", operatorList)

	var expressions []string       // 产生的表达式
	var resultExpressions []string // 能够满足number points 的表达式
	for i := 0; i < len(*resultList); i++ {
		for j := 0; j < len(*operatorList); j++ {

			// 第1种情况 (a+b) + (c+d)
			ex1 := "(" + strconv.Itoa((*resultList)[i][0]) + getOperator((*operatorList)[j][0]) + strconv.Itoa((*resultList)[i][0]) + ")" +
				getOperator((*operatorList)[j][1]) + "(" + strconv.Itoa((*resultList)[i][2]) +
				getOperator((*operatorList)[j][2]) + strconv.Itoa((*resultList)[i][3]) + ")"
			expressions = append(expressions, ex1)
			if Calculate(ex1) == number {
				resultExpressions = append(resultExpressions, ex1)
			}

			// 第2种情况 ((a+b)+c)+d
			ex2 := "(" + "(" + strconv.Itoa((*resultList)[i][0]) + getOperator((*operatorList)[j][0]) + strconv.Itoa((*resultList)[i][0]) + ")" +
				getOperator((*operatorList)[j][1]) + strconv.Itoa((*resultList)[i][2]) + ")" +
				getOperator((*operatorList)[j][2]) + strconv.Itoa((*resultList)[i][3])
			expressions = append(expressions, ex2)
			if Calculate(ex2) == number {
				resultExpressions = append(resultExpressions, ex2)
			}

			// 第3种情况 (a +(b+c)) + d
			ex3 := "(" + strconv.Itoa((*resultList)[i][0]) + getOperator((*operatorList)[j][0]) + "(" + strconv.Itoa((*resultList)[i][0]) + getOperator((*operatorList)[j][1]) + strconv.Itoa((*resultList)[i][2]) + ")" +
				")" +
				getOperator((*operatorList)[j][2]) + strconv.Itoa((*resultList)[i][3])
			expressions = append(expressions, ex3)
			if Calculate(ex3) == number {
				resultExpressions = append(resultExpressions, ex3)
			}
		}

	}
	log.Println("resultExpressions:", resultExpressions)
}

// Perm return numbers permutation from its left to right
func Perm(numbers []int, left, right int, result *NumberList) {
	if left == right {
		temp := new([]int)
		for i := 0; i < len(numbers); i++ {
			*temp = append(*temp, numbers[i])
		}
		*result = append(*result, *temp)
	} else {
		for i := left; i <= right; i++ {
			numbers[left], numbers[i] = numbers[i], numbers[left]
			Perm(numbers, left+1, right, result)
			numbers[left], numbers[i] = numbers[i], numbers[left]
		}
	}

}

// NumberList store the all numbers permutation
type NumberList [][]int

func getOperator(o int) string {
	switch o {
	case PLUS:
		return "+"
	case SUBSTRACT:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	}

	return "+"
}
