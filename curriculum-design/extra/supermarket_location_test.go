/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : supermarket_location_test.go
#   Last Modified : 2018-12-27 14:16
#   Describe      :超市选址 附加题目 3 带权有向图中心 测试
#
# ====================================================*/

package extra

import (
	"log"
	"testing"
)

func TestFindSupermarketLocation(t *testing.T) {
	matrix := new([][]int)
	pNumber := new([]int)

	p1 := [...]int{0, 9999, 3, 9999}
	p2 := [...]int{2, 0, 9999, 9999}
	p3 := [...]int{9999, 7, 0, 1}
	p4 := [...]int{6, 9999, 9999, 0}

	*matrix = append(*matrix, p1[:])
	*matrix = append(*matrix, p2[:])
	*matrix = append(*matrix, p3[:])
	*matrix = append(*matrix, p4[:])

	*pNumber = append(*pNumber, 20)
	*pNumber = append(*pNumber, 30)
	*pNumber = append(*pNumber, 40)
	*pNumber = append(*pNumber, 25)

	log.Println("地址信息如下:")
	for _, v := range *matrix {
		log.Println(v)
	}

	position, maxLength, path := FindSupermarketLocation(*matrix, *pNumber)
	log.Println("超市应当建立在:", position)
	log.Println("超市到各个地方总距离为:", maxLength)
	log.Println("超市路径为:", path)

}
