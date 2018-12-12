package bubblesort

// BubbleSort 使用冒泡排序
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
