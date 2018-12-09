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
