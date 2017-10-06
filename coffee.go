package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
	"net"
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

func waitForInternetConnection(try int) {
	if try <= 0 {
		log.Panic("No internet connection detected")
	}
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Println("Waiting for connection. Try number " + string(try))
		time.Sleep(time.Second * 5)
		waitForInternetConnection(try - 1)
	}
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

	waitForInternetConnection(10)
	
	t, err := hc.NewIPTransport(hc.Config{Pin: "84297450"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}