package commands

import "time"

type About struct {
	Name             string
	Description      string
	Cooldown         int
	Alias            []int
	AdditionalFields struct {
		Author  string
		Version int
		License string
		Uptime  time.Duration
		Name    string
	}
}
