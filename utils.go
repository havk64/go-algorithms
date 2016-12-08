package utils

import (
	"math/rand"
	"time"
)

func RandArray(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
