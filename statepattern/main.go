package main

import "fmt"

//state

type Salary interface {
	state()
}

// concrete implementation

type credited struct{}

func (c *credited) state() {
	fmt.Println("Salary is credited")
}

type notcredited struct{}

func (nc *notcredited) state() {
	fmt.Println("Salary is not credited")
}

type statecontext struct {
	currentstatus Salary
}

func getcontext() *statecontext {
	return &statecontext{
		currentstatus: &notcredited{},
	}
}

func (sc *statecontext) setstate(state Salary) {
	sc.currentstatus = state
}

func (sc *statecontext) getstate() {
	sc.currentstatus.state()
}

// this is the client

func main() {
	currentsalary := getcontext()
	currentsalary.getstate()
	currentsalary.setstate(&credited{})
	currentsalary.getstate()

}
