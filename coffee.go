package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"log"
	"time"
	"net"
)

func waitForInternetConnection(try int) {
	if try <= 0 {
		log.Panic("No internet connection detected")
	}
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Printf("Waiting for connection. Try number %d", try)
		time.Sleep(time.Second * 5)
		waitForInternetConnection(try - 1)
	}
}

type RemoteEventHandler interface {
	RemoteValueUpdated(on bool)
}

func main() {
	// Start GPIO
	handler := GPIOHandler{}
	handler.SetupGPIO()

	info := accessory.Info{
		Name: "Raspi Coffee",
		Manufacturer: "macostea",
	}

	log.Println(info)

	acc := accessory.NewOutlet(info)

	acc.Outlet.On.OnValueRemoteUpdate(func (on bool) {
		handler.RemoteValueUpdated(on)
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