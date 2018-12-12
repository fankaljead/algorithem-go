package utils

import (
	"math/rand"
)

// GetArrayBySize return a int array use size
func GetArrayBySize(size int) []int {
	numbers := make([]int, 0)

	for i := 0; i < size; i++ {
		num := rand.Intn(1000)
		numbers = append(numbers, num)
	}

	return numbers
}
