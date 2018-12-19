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

var (
	bottles     Bottles
	bottlesList []Bottles
	root        Node
	last        Node
)

func init() {
	bottles = append(bottles, Bottle{8, 8})
	bottles = append(bottles, Bottle{5, 0})
	bottles = append(bottles, Bottle{3, 0})
	bottlesList = append(bottlesList, bottles)

}

// FindFourPints is to find 4 pints
func FindFourPints() {

	tempBottles := bottles.Copy()
	root.Bottles = bottles
	current := root
	last = root

	for _, v := range current.Son {
		temp := v.Bottles.Copy()
		for _, v1 := range temp {
			for _, v2 := range temp {
				if v1 != v2 {
					v1.PourInto(&v2)
				}
			}
		}
	}

	for {
		tempBottles = tempBottles.Copy()

		for _, v1 := range tempBottles {
			for _, v2 := range tempBottles {
				if v1 != v2 {
					v1.PourInto(&v2)
				}
			}
		}

		bottlesList = append(bottlesList, tempBottles)

		if tempBottles.Has(4) {
			break
		}
	}

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

// PourInto cap water to another bottle
func (bs *Bottle) PourInto(abs *Bottle) bool {
	cap := abs.Capacity - abs.Used

	if cap == 0 {
		return false
	}

	if bs.Used <= 0 {
		return false
	}

	bs.Used -= cap
	abs.Used += cap
	return true

}

// Copy can copy bottles return a new bottles
func (bs Bottles) Copy() (nbs Bottles) {
	for _, v := range bs {
		nbs = append(nbs, Bottle{Capacity: v.Capacity, Used: v.Used})
	}
	return
}

// Node store a node for Bottles Tree
type Node struct {
	Bottles Bottles
	Parent  *Node
	Son     []Node
}

// BottlesTree store
type BottlesTree struct {
	Root *Node
}
