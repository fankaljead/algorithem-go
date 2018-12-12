package bubblesort

import (
	"gitlab.com/fankaljead/algorithem-go/chapter_3_brute_force/utils"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	a := utils.GetArrayBySize(1000)
	BubbleSort(a[:])
	for i := 0; i < len(a)-2; i++ {
		if a[i] > a[i+1] {
			t.Error()
		}
	}
}
