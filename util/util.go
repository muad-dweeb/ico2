package util

import (
	"math/rand"
	"time"
)

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
