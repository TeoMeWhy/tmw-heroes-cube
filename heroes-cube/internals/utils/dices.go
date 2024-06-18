package utils

import (
	"math/rand"
	"time"
)

var source = rand.New(rand.NewSource(time.Now().Unix()))

func DrawDM(m, n int) []int {

	numbers := []int{}

	for i := 1; i <= n; i++ {
		number := source.Intn(m) + 1
		numbers = append(numbers, number)
	}
	return numbers
}
