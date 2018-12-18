/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : eight_pints.go
#   Last Modified : 2018-12-18 22:42
#   Describe      : 遍历第2题 8品脱和5品脱和3品脱问题
#
# ====================================================*/

package traverse

import (
	"log"
)

// FindFourPints is to find 4 pints
func FindFourPints() {
	var bottles Bottles
	bottles = append(bottles, Bottle{8, 8})
	bottles = append(bottles, Bottle{5, 0})
	bottles = append(bottles, Bottle{3, 0})

	log.Println(bottles)
}

// Bottles can store some bottles
type Bottles []Bottle

// Bottle store the total capacity and the used capacity of a bottle
type Bottle struct {
	Capacity, Used int
}

// Has returns bottles has u used or not
func (bs Bottles) Has(u int) bool {
	for _, v := range bs {
		if v.Used == u {
			return true
		}
	}

	return false
}
