package selectsort

import (
	"fmt"
	"gitlab.com/fankaljead/algorithem-go/chapter_3_brute_force/utils"
	"testing"
)

// TestSelectSort for Select Sort
func TestSelectSort(t *testing.T) {
	a := utils.GetArrayBySize(10)
	SelectSort(a[:])
	for i := 0; i < len(a)-2; i++ {
		if a[i] > a[i+1] {
			t.Error()
		}
		fmt.Printf("%d\t", a[i])
	}
	fmt.Println()
}
