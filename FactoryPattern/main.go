package main

import "fmt"

// App

type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Honda City"
}

type HatchBackType struct {
	Name string
}

func getNewHatchBack() *HatchBackType {
	return &HatchBackType{}
}

func (h *HatchBackType) getCar() string {
	return "Tata Tiago"
}

// client

func main() {
	getCarFactory(1)
	getCarFactory(2)
}

// Factory class

func getCarFactory(cartype int) {
	var car Car
	if cartype == 1 {
		car = getNewHatchBack()
	} else {
		car = getNewSedan()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)
}
