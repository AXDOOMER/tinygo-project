package main

import (
	"machine"
	"time"
)

func main() {
	go motor()
	pump()
}

func motor() {
	led := machine.D7
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(500)
		led.Low()
		delay(500)
	}
}

func pump() {
	led := machine.D8
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(1000)
		led.Low()
		delay(1000)
	}
}

func delay(t int64) {
	time.Sleep(time.Duration(1000000 * t))
}
