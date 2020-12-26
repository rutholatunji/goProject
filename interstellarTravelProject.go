package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println(fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	if planet == "Venus" {
		fuel = 300000
		fmt.Println("Venus costs 300000 pounds of fuel")
	} else if planet == "Mecury" {
		fuel = 500000
		fmt.Println("Mecury costs 500000 pounds of fuel")
	} else if planet == "Mars" {
		fuel = 700000
		fmt.Println("Mars costs 700000 pounds of fuel")
	} else {
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Greetings earthlings, welcome to", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there")
}

func flytoPlanet(planet string, fuel int) int {
	var fuelRemaining int
	var fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining > fuelCost {
		greetPlanet(planet)
		fmt.Println(fuelRemaining-fuelCost, "pounds of fuel remaining!")
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	fmt.Println("Welcome to Interstellar airlines!")

	fuel := 1000000
	planetChoice := "Venus"

	fuel = flytoPlanet(planetChoice, fuel)
}
