package main

import (
	"fmt"
	"ico2/util"
)

func main() {

	n := util.GetRandomInt()
	fmt.Printf("Random Number: %d\n", n)

	h, m, s, ms := util.MillisToReadable()
	fmt.Printf("Hour: %d, Minute: %d, Second: %d, Millisecond: %d\n", h, m, s, ms)
}
