package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

type GPIOHandler struct {
	controlPin rpio.Pin
}

func (h *GPIOHandler) RemoteValueUpdated(on bool) {
	h.controlPin.High()
	time.Sleep(time.Second * 1)
	h.controlPin.Low()
}

func (h *GPIOHandler) SetupGPIO() {
	err := rpio.Open()
	if err != nil {
		log.Fatal("Cannot start GPIO module")
	}

	h.controlPin = rpio.Pin(23)
	h.controlPin.Output()
}
