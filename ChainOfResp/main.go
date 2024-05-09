package main

import "fmt"

type step interface {
	execute()     // Function to be executed next.
	setNext(step) // Function to set the next func to be executed.
}

type openlap struct {
	next step
}

type pressStartButton struct {
	next step
}

type enterPassword struct {
	next step
}

// Open lap

func (r *openlap) execute() {
	fmt.Println("Opening Laptop Flap")
	r.next.execute()
}

func (r *openlap) setNext(next step) {
	r.next = next
}

// Press Button Below

func (d *pressStartButton) execute() {
	fmt.Println("Pressing start button")
	d.next.execute()
}

func (d *pressStartButton) setNext(next step) {
	d.next = next
}

// Enter Password

func (m *enterPassword) execute() {
	fmt.Println("Entering password for Laptop")
	fmt.Println("Laptop Opened!")
}

func (m *enterPassword) setNext(next step) {
	m.next = next
}

func main() {
	//Set next for enterPassword step
	enterPassword := &enterPassword{}

	//Set next for pressStartButton step
	pressStartButton := &pressStartButton{}
	pressStartButton.setNext(enterPassword)

	//Set next for openFlap step
	openlap := &openlap{}
	openlap.setNext(pressStartButton)

	openlap.execute()
}
