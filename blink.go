package main

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	pin := gpioreg.ByName("GPIO17")

	// Forever...
	for {
		pin.Out(gpio.High) // turn LED on

		time.Sleep(time.Second) // wait one second

		pin.Out(gpio.Low) // turn LED off
	}
}
