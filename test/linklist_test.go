/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : linklist_test.go
#   Last Modified : 2018-12-12 12:47
#   Describe      :
#
# ====================================================*/

package list

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var list LinkedList
	list.Add(2)
	list.Add(3)
	list.Add(4)

	// fmt.Printf("%v\n", list)
	fmt.Println(list.GetFirst())
	fmt.Println(list.Get(1))
	fmt.Println(list)
}
