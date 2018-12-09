package selectsort

import (
	"fmt"
	"github.com/fankaljead/algorithm/chapter_3_brute_force/utils"
	"testing"
)

// TestSelectSort for Select Sort
func TestSelectSort(t *testing.T) {
	a := util.GetArrayBySize(1000)
	SelectSort(a[:])
	for i := 0; i < len(a)-2; i++ {
		if a[i] > a[i+1] {
			fmt.Println(a)
			t.Error()
		}
	}
}
