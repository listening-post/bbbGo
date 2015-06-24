package main

import (
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/bbb" // This loads the RPi driver
)

func main() {
	for {
		embd.LEDToggle("USR3")
		time.Sleep(2000* time.Millisecond)
	}
}
