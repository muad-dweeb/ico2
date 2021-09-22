package main

import (
	"fmt"
	"ico2/util"
	"time"
)

func applicationLaunch() time.Time {
	return time.Now()
}

func main() {
	// returns a random int, random int is seeded
	// by time.Now().UnixNano()
	n := util.GetRandomInt()

	// test need to eventually remove as it's not incredibly useful.
	// prints random number so we can see it though.
	fmt.Printf("Random Number: %d\n", n)

	// launches application - as of now only return the time the function ran
	//TODO: - add debug flag.
	t := applicationLaunch()

	// test to show uptime actually works
	// time.Sleep(90 * time.Second)
	uptime := util.UptimeCount(t)

	// also eventually remove, used as a proof of concept to show that uptime
	// is actually displaying and being calculated correctly.
	fmt.Printf("Uptime is: %v\n", uptime)
}
