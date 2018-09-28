package main

import (
	"fmt"
	"math"
)

type setup struct {
	farmers int // the amont of farmers in the system
	fish    int // the amount of fish in the pond to begin with
	newFish int // the amount of new fishes to be added to the system at the end of a day.
}

type farmer struct {
	honest bool // tells if the farmer is being honest (choosing to abide by the rules)
	amount int  // the amount of fish the farmer takes from the pond.
}

func main() {
	welcome()
	systemSetup := setup{4, 11, 1}
	fmt.Println("Initial conditions ", systemSetup.getInitialConditions())
	systemSetup.createNewFish()
	systemSetup.remainingFish()
	fmt.Println("farmers ", systemSetup.createFarmer())
}

// init method sets the initial condition of the experiment
func (f farmer) details() string {
	return fmt.Sprintf("Honest --> %t | Taking --> %d fish", f.honest, f.amount)
}
func (s setup) amount() int {
	return s.farmers
}

// The goFishing function start the simulatoin
func goFishing(f farmer) {

}

// getInitialConditions returns the starting conditions of the system.
func (s setup) getInitialConditions() string {
	return fmt.Sprintf("There are %d farmers in the system with %d fish to begin with. \n%d new fish to be added at the end of the day.", s.farmers, s.fish, s.newFish)
}

func (s setup) remainingFish() {
	fmt.Printf("\n%d fish remaining", s.fish)
}

// createFarmer method takes the number of farmers in the systen and creates individual farmers
func (s setup) createFarmer() map[int]interface{} {
	var m = make(map[int]interface{})
	for i := 1; i < s.amount(); i++ {
		m[i] = farmer{true, 1}
	}
	return m
}

// t
func (s *setup) createNewFish() {
	fmt.Println("Re-populating pond...")
	newOnes := math.Floor(float64(s.fish / 2))
	s.fish = s.fish + int(newOnes)
	fmt.Printf("\n%d new fish. Total now at %d.", int(newOnes), s.fish)
}

// welcome message to be displayed to the user
func welcome() {
	fmt.Println("This is a simulatoin of the tragedy of the commons, a famous thought experiment which investigates what happens when individuals choose to go against what is good for the collective. ")
}
