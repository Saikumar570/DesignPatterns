package main

import "fmt"

// Target

type mobile interface {
	chargeAppleMobile()
}

// Concrete Prototype Implementation
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging APPLE mobile")
}

// Adaptee
type android struct{}

func (a *android) chargeAndroidMobile() {
	fmt.Println("Charging ANDROID mobile")
}

// Adapter
type androiddapter struct {
	android *android
}

func (ad *androiddapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
}

// Client
type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}

func main() {
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	android := &android{}
	androiddapter := &androiddapter{
		android: android,
	}
	client.chargeMobile(androiddapter)
}
