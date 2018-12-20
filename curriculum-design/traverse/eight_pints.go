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

	root.Bottles = bottles
	current := root

	preBottleListSize := len(bottlesList)
	var sons [][]Node
	sons = append(sons, append(current.Son, current))

	for {
		for _, v1 := range current.Bottles {
			for _, v2 := range current.Bottles {
				if !v1.Equal(v2) {
					tempV1 := v1.Copy()
					tempV2 := v2.Copy()
					if tempV1.PourInto(&tempV2) {
						log.Println(tempV1.Capacity, " pour into ", tempV2.Capacity)
						var bottlesNode Bottles
						bottlesNode = append(bottlesNode, tempV1)
						bottlesNode = append(bottlesNode, tempV2)

						// 创建新的Bottles
						// Create new Bottles after pouring
						for _, v3 := range current.Bottles {
							if !v3.Equal(tempV1) && !v3.Equal(tempV2) {
								tempV3 := v3.Copy()
								bottlesNode = append(bottlesNode, tempV3)
							}
						}
						if bottlesNode.Has(4) {
							return
						}

						// 判断节点是否存在
						// judge the new bottles is existed or not after pouring
						flag := true
						for _, alreadyExist := range bottlesList {
							if alreadyExist.Equal(bottlesNode) {
								flag = false
								break
							}
						}
						// 没有满足结束条件，但未出现过，将其加入 current 的子节点中
						if flag {
							var tempNode Node
							tempNode.Bottles = bottlesNode
							tempNode.Parent = &current
							current.Son = append(current.Son, tempNode)
							bottlesList = append(bottlesList, bottlesNode)

						}

					}
				}
			}
		}

		// 判断终止，当长度不在变换，说明已经遍历出所以情况
		if preBottleListSize == len(bottlesList) {
			log.Println("no answer")
			return
		}
	}

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
func (b *Bottle) PourInto(ab *Bottle) bool {
	cap := ab.Capacity - ab.Used

	if cap == 0 {
		return false
	}

	if b.Used <= 0 {
		return false
	}

	b.Used -= cap
	ab.Used += cap
	return true

}

// Equal judge two bottle is equal or not
func (b Bottle) Equal(ab Bottle) bool {
	if b.Capacity == ab.Capacity && b.Used == ab.Used {
		return true
	}
	return false
}

// Equal judge two bottle is equal or not
func (bs Bottles) Equal(abs Bottles) bool {
	for _, v1 := range bs {
		for _, v2 := range abs {
			if v1.Equal(v2) {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

// Copy can copy bottles return a new bottles
func (bs Bottles) Copy() (nbs Bottles) {
	for _, v := range bs {
		nbs = append(nbs, Bottle{Capacity: v.Capacity, Used: v.Used})
	}
	return
}

// Copy can copy bottle return a new bottle
func (b *Bottle) Copy() (nb Bottle) {
	nb.Capacity = b.Capacity
	nb.Used = b.Used

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
