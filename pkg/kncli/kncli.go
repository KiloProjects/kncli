package kncli

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	// Initialize the random number generator with a unique seed based on current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(101)
}
