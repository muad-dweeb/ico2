package util

import (
	"log"
	"math/rand"
	"time"
)

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func RollOutput() {

}

func UptimeCount(t time.Time) time.Duration {
	dur, err := time.ParseDuration(time.Duration.String(time.Since(t)))
	if err != nil {
		log.Fatal(err)
	}
	return dur
}
