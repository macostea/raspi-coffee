package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

func togglePin(pin rpio.Pin) {
	pin.High()
	time.Sleep(time.Second * 1)
	pin.Low()
}

func turnCoffeeMakerOn(pin rpio.Pin) {
	log.Println("Turn Coffee Maker On")
	togglePin(pin)
}

func turnCoffeeMakerOff(pin rpio.Pin) {
	log.Println("Turn Coffee Maker Off")
	togglePin(pin)
}

func main() {
	// Start GPIO
	err := rpio.Open()
	if err != nil {
		log.Fatal("Cannot start GPIO module")
	}

	pin := rpio.Pin(23)
	pin.Output()

	info := accessory.Info{
		Name: "Raspi Coffee",
		Manufacturer: "macostea",
	}

	log.Println(info)

	acc := accessory.NewOutlet(info)

	acc.Outlet.On.OnValueRemoteUpdate(func (on bool) {
		if on {
			turnCoffeeMakerOn(pin)
		} else {
			turnCoffeeMakerOff(pin)
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "84297450"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}