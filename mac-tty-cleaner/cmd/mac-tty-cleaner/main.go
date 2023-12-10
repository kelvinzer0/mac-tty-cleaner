package main

import (
	"log"

	"github.com/kelvinzer0/mac-tty-cleaner/internal/tty"
)

func main() {
	// Specify the command you want to run (e.g., "clear")
	command := "clear"

	// Get a list of TTY devices
	ttyDevices, err := tty.GetTTYDevices()
	if err != nil {
		log.Printf("Error getting TTY devices: %v\n", err)
		return
	}

	// Iterate over TTY devices and run the command on each one
	for _, ttyDevice := range ttyDevices {
		err := tty.RunCommandOnTTY(command, ttyDevice)
		if err != nil {
			log.Printf("Error running command on %s: %v\n", ttyDevice, err)
		} else {
			log.Printf("Command executed successfully on %s\n", ttyDevice)
		}
	}
}
