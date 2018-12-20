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
	nodeLists   Nodes
	resultList  []Bottles
)

func init() {
	bottles = append(bottles, Bottle{112, 99})
	bottles = append(bottles, Bottle{50, 25})
	bottles = append(bottles, Bottle{12, 1})
	bottles = append(bottles, Bottle{2, 0})
	bottles = append(bottles, Bottle{5, 4})
	// bottles = append(bottles, Bottle{12, 11})
	// bottles = append(bottles, Bottle{5, 0})
	// bottles = append(bottles, Bottle{3, 0})
	// bottles = append(bottles, Bottle{4, 0})
	bottlesList = append(bottlesList, bottles)

}

// PourDo do pour
func PourDo() {
	var nodes Nodes
	nodes = append(nodes, Node{Bottles: bottles, Parent: &root})
	resultNode := Find(nodes, 40)
	current := &resultNode

	for current.Parent != nil {
		resultList = append(resultList, current.Bottles)
		current = current.Parent
	}

	for i := len(resultList) - 1; i >= 0; i-- {
		log.Printf("bottles:%v\n", resultList[i])
	}
}

// Find is to find 4 pints
func Find(nodes Nodes, number int) (node Node) {
	// log.Println("nodes:", len(nodes))
	if len(nodes) == 0 {
		// node = Node{Bottles: bottles}
		node = Node{}
		return
	}
	cloneNode := nodes.Clone()
	var newNodes Nodes
	for i := 0; i < len(cloneNode); i++ {
		bottleNode := cloneNode[i]
		// log.Printf("bottleNode address:%p\n", &bottleNode)
		// for _, bottleNode := range cloneNode {
		for bottle1 := 0; bottle1 < len(bottleNode.Bottles); bottle1++ {
			for bottle2 := 0; bottle2 < len(bottleNode.Bottles); bottle2++ {
				// for bottle1, _ := range bottleNode.Bottles {
				//     for bottle2, _ := range bottleNode.Bottles {
				if !bottleNode.Bottles[bottle1].Equal(bottleNode.Bottles[bottle2]) {
					newNode, succeed := Pour(bottle1, bottle2, bottleNode)

					// log.Println("newNode:", newNode)
					if succeed {
						if !nodeLists.IsIn(&newNode) {
							// log.Println("succeed")
							newNode.Parent = &bottleNode
							nodeLists = append(nodeLists, newNode)
							newNodes = append(newNodes, newNode)
							// log.Println("append newNode:", newNode)
							// log.Printf("bottle node:%v\n", bottleNode)
							// log.Println("current node:", newNode)
						}
						if newNode.Bottles.Has(number) {
							node = newNode
							// log.Println("return 4 node:", node)
							return
						}

					}
				}

			}

		}
	}
	// log.Println("newNodes:", newNodes)
	return Find(newNodes, number)
}

// Pour into bottle2 from bottle1 in node.Bottles
func Pour(bottle1, bottle2 int, node Node) (newNode Node, can bool) {
	newNode = node.Clone()
	// log.Println("pour node:", node)
	can = newNode.Bottles[bottle1].PourInto(&newNode.Bottles[bottle2])
	return

}

// // FindFourPints is to find 4 pints
// func FindFourPints() {
//
//     root.Bottles = bottles
//     current := root
//
//     preBottleListSize := len(bottlesList)
//     var sons [][]Node
//     sons = append(sons, append(current.Son, current))
//
//     for {
//         for _, v1 := range current.Bottles {
//             for _, v2 := range current.Bottles {
//                 if !v1.Equal(v2) {
//                     tempV1 := v1.Copy()
//                     tempV2 := v2.Copy()
//                     if tempV1.PourInto(&tempV2) {
//                         log.Println(tempV1.Capacity, " pour into ", tempV2.Capacity)
//                         var bottlesNode Bottles
//                         bottlesNode = append(bottlesNode, tempV1)
//                         bottlesNode = append(bottlesNode, tempV2)
//
//                         // 创建新的Bottles
//                         // Create new Bottles after pouring
//                         for _, v3 := range current.Bottles {
//                             if !v3.Equal(tempV1) && !v3.Equal(tempV2) {
//                                 tempV3 := v3.Copy()
//                                 bottlesNode = append(bottlesNode, tempV3)
//                             }
//                         }
//                         if bottlesNode.Has(4) {
//                             return
//                         }
//
//                         // 判断节点是否存在
//                         // judge the new bottles is existed or not after pouring
//                         flag := true
//                         for _, alreadyExist := range bottlesList {
//                             if alreadyExist.Equal(bottlesNode) {
//                                 flag = false
//                                 break
//                             }
//                         }
//                         // 没有满足结束条件，但未出现过，将其加入 current 的子节点中
//                         if flag {
//                             var tempNode Node
//                             tempNode.Bottles = bottlesNode
//                             tempNode.Parent = &current
//                             current.Son = append(current.Son, tempNode)
//                             bottlesList = append(bottlesList, bottlesNode)
//
//                         }
//
//                     }
//                 }
//             }
//         }
//
//         // 判断终止，当长度不在变换，说明已经遍历出所以情况
//         if preBottleListSize == len(bottlesList) {
//             log.Println("no answer")
//             return
//         }
//     }
//
// }

// Bottles can store some bottles
type Bottles []Bottle

// Bottle store the total capacity and the used capacity of a bottle
type Bottle struct {
	Capacity, Used int
}

// Has returns bottles has u used or not
func (bs *Bottles) Has(u int) bool {
	// log.Println("bs:", bs)
	for _, v := range *bs {
		if v.Used == u {
			return true
		}
	}

	return false
}

// PourInto cap water to another bottle
func (b *Bottle) PourInto(ab *Bottle) bool {
	cap := ab.Capacity - ab.Used

	if cap <= 0 {
		return false
	}

	if b.Used <= 0 {
		return false
	}

	if b.Used < cap {
		cap = b.Used
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
	for i := 0; i < len(bs); i++ {
		if !bs[i].Equal(abs[i]) {
			return false
		}
	}
	// for _, v1 := range bs {
	//     for _, v2 := range abs {
	//         if v1.Equal(v2) {
	//             continue
	//         } else {
	//             return false
	//         }
	//     }
	// }

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
}

// Clone return a nodes copy
func (node *Node) Clone() (newNode Node) {
	newNode.Parent = node.Parent
	newNode.Bottles = node.Bottles.Copy()
	return
}

// Equal return node is equal to n or not
func (node *Node) Equal(n *Node) bool {
	if node.Bottles.Equal(n.Bottles) {
		return true
	}

	return false
}

// Nodes store some nodes
type Nodes []Node

// Clone return a nodes copy
func (nodes *Nodes) Clone() (newNodes Nodes) {
	for _, v := range *nodes {
		node := v.Clone()
		newNodes = append(newNodes, node)
	}
	return
}

// IsIn return the node is in the nodeslist or not
func (nodes *Nodes) IsIn(node *Node) bool {

	// log.Println("isin node:", node)
	for _, v := range *nodes {
		if v.Equal(node) {
			return true
		}
	}

	return false
}

// BottlesTree store
type BottlesTree struct {
	Root *Node
}
