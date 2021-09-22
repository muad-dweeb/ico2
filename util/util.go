package util

import (
	"math/rand"
	"time"
)

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func MillisToReadable() (int, int, int, int64) {
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()
	ms := time.Now().UnixMilli()
	return h, m, s, ms

}
