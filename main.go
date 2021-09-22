package main

import (
	"fmt"
	"ico2/util"
)

func main() {

	n := util.GetRandomInt()
	fmt.Printf("Random Number: %d\n", n)
	t := util.ApplicationLaunch()
	// test to show uptime actually works
	// time.Sleep(90 * time.Second)
	uptime := util.UptimeCount(t)

	fmt.Printf("Uptime is: %v\n", uptime)
}
