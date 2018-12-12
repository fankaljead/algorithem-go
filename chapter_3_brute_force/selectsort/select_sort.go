/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : select_sort.go
#   Last Modified : 2018-12-09 22:00
#   Describe      : 蛮力法选择排序
#
# ====================================================*/

package selectsort

// SelectSort 对一个数组进行选择排序
func SelectSort(a []int) {
	for i := 0; i < len(a)-2; i++ {
		min := i
		for j := i + 1; j < len(a)-1; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		// swap a[i] and a[min]
		a[i], a[min] = a[min], a[i]
	}
}
