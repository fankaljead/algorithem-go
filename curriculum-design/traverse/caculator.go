/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : caculator.go
#   Last Modified : 2018-12-20 22:49
#   Describe      : 计算表达式
#
# ====================================================*/

package traverse

import (
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// LEVEL1 + - 的先后顺序
	LEVEL1 = iota
	// LEVEL2 + - 的先后顺序
	LEVEL2
	// LEVEL3 + - 的先后顺序
	LEVEL3
)

func init() {
	po["+"] = LEVEL1
	po["-"] = LEVEL1
	po["*"] = LEVEL2
	po["/"] = LEVEL2
}

var (
	po = make(map[string]int)
)

func parseExp(s string) (exps []string, err error) {
	re, err := regexp.Compile("[0-9]+|[+*/\\-\\(\\)]{1}")
	if err != nil {
		fmt.Println("regexp compile error:", err)
		return
	}
	for _, exp := range re.FindAll([]byte(s), -1) {
		exps = append(exps, string(exp))
	}
	return
}

func isPop(list *list.List, s string) (op []string, ok bool) {
	switch string(s) {
	case "(":
		ok = false
		return
	case ")":
		ok = true
		cur := list.Back()
		for {
			prev := cur.Prev()
			if curValue, ok2 := cur.Value.(string); ok2 {
				if string(curValue) == "(" {
					list.Remove(cur)
					return
				}
				op = append(op, curValue)
				list.Remove(cur)
				cur = prev
			}
		}
	default:
		for cur := list.Back(); cur != nil; {
			prev := cur.Prev()
			if curValue, ok2 := cur.Value.(string); ok2 {
				if level, ok3 := po[curValue]; ok3 && level >= po[s] {
					ok = true
					op = append(op, curValue)
					// fmt.Println(op)
					list.Remove(cur)
				} else if curValue == "(" {
					// fmt.Println(curValue, op)
					if len(op) != 0 {
						ok = true
					} else {
						ok = false
					}
					return
				}
			}
			cur = prev
		}
	}
	return
}

func isOperate(s string) bool {
	re, _ := regexp.Compile("[+*/\\-\\(\\)]{1}")
	ok := re.MatchString(s)
	// fmt.Println(ok, s)
	return ok
}

func pre2stuf(exps []string) (exps2 []string) {
	list1 := list.New()
	list2 := list.New()

	for _, exp := range exps {
		if isOperate(exp) {
			if op, ok := isPop(list1, exp); ok {
				for _, s := range op {
					list2.PushBack(s)
				}
			}
			if exp == ")" {
				continue
			}
			list1.PushBack(exp)
		} else {
			list2.PushBack(exp)
		}
	}

	for cur := list1.Back(); cur != nil; cur = cur.Prev() {
		// fmt.Print(cur.Value)
		list2.PushBack(cur.Value)
	}

	for cur := list2.Front(); cur != nil; cur = cur.Next() {
		if curValue, ok := cur.Value.(string); ok {
			exps2 = append(exps2, curValue)
		}
	}
	return
}

func calculate(exps []string) int {
	list1 := list.New()

	for _, s := range exps {
		if isOperate(s) {
			back := list1.Back()
			prev := back.Prev()
			backVal, _ := back.Value.(int)
			prevVal, _ := prev.Value.(int)
			var res int
			switch s {
			case "+":
				res = prevVal + backVal
			case "-":
				res = prevVal - backVal
			case "*":
				res = prevVal * backVal
			case "/":
				res = prevVal / backVal
			}
			list1.Remove(back)
			list1.Remove(prev)
			list1.PushBack(res)
		} else {
			v, _ := strconv.Atoi(s)
			list1.PushBack(v)
		}
	}
	if list1.Len() != 1 {
		fmt.Println("caculate error")
		os.Exit(1)
	}
	res, _ := list1.Back().Value.(int)
	return res
}

// func main() {
//     po["+"] = LEVEL1
//     po["-"] = LEVEL1
//     po["*"] = LEVEL2
//     po["/"] = LEVEL2
//
//     // a := "1+23*1/123+(4 + 2 *5)"
//     fmt.Println("Please input express:")
//     var a string
//     fmt.Scanln(&a)
//     fmt.Println(a)
//
//     a = strings.Replace(a, " ", "", -1)
//     exps, err := parseExp(a)
//     if err != nil {
//         os.Exit(1)
//     }
//     fmt.Println(exps)
//     exps2 := pre2stuf(exps)
//     fmt.Println(exps2)
//     fmt.Println(caculate(exps2))
// }

// Calculate a expression
func Calculate(expression string) int {
	expression = strings.Replace(expression, " ", "", -1)
	exps, err := parseExp(expression)
	if err != nil {
		os.Exit(1)
	}

	exps2 := pre2stuf(exps)

	return calculate(exps2)
}
