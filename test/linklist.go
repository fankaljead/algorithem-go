/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : linklist.go
#   Last Modified : 2018-12-12 12:38
#   Describe      :
#
# ====================================================*/

package list

import (
	"bytes"
	// "fmt"
)

// Node do
type Node struct {
	Next    *Node
	element interface{}
}

// LinkedList do
type LinkedList struct {
	Head *Node
	Size int
}

// Add a element to linked list
func (list *LinkedList) Add(e interface{}) {
	temp := list.Head

	var node Node
	node.element = e
	if list.Size == 0 {
		list.Head = &node
	} else {
		for i := 0; i < list.Size-1; i++ {
			temp = temp.Next
		}
		temp.Next = &node
	}
	list.Size++
}

// String return linklist elements
func (list *LinkedList) String() string {
	var buf bytes.Buffer
	temp := list.Head
	for i := 0; i < list.Size; i++ {
		// s += temp.element.(string)
		buf.WriteString(temp.element.(string))
		// fmt.Println(temp.element)
		temp = temp.Next
	}

	return buf.String()
}

// GetFirst return the first element of linklist
func (list *LinkedList) GetFirst() interface{} {
	return list.Head.element
}

// Get return the index element of linklist
func (list *LinkedList) Get(index int) interface{} {
	temp := list.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.element
}

func (node *Node) String() string {
	return node.element.(string)
}
