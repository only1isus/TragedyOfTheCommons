package main

import (
	"fmt"
	"log"
	"math"

	"github.com/fatih/color"
)

type setup struct {
	farmers int // the amont of farmers in the system
	fish    int // the amount of fish in the pond to begin with
	newFish int // the amount of new fishes to be added to the system at the end of a day.
	period  int // tracks the of 'days'
}

type farmer struct {
	honest bool // tells if the farmer is being honest (choosing to abide by the rules) a farmer is 'honest' if he only takes one fish
	amount int  // the amount of fish the farmer takes from the pond.
}

func main() {
	// Welcome()
	systemSetup := setup{4, 12, 1, 0} // this is where the starting conditions are set. farmers, fish, amount taken
	display := color.New(color.FgGreen, color.Bold)
	farmers := systemSetup.createFarmer()
	display.Printf("%d farmers, %d fish available, %d days\n", systemSetup.amountOfFarmers(), systemSetup.amountOfFish(), systemSetup.period)
	GoFishing(farmers, &systemSetup)
}

// details method tells the initial condition of the experiment
func (f farmer) details() string {
	return fmt.Sprintf("Honest --> %t | Taking --> %d fish", f.honest, f.amount)
}

// this method returns the amount of farmers in the system
func (s *setup) amountOfFarmers() int {
	return s.farmers
}

// this method returns the amount of fish in the system
func (s *setup) amountOfFish() int {
	return s.fish
}

// this method updates the amount of fish in the pond
func (s *setup) updateFish(amount int) {
	s.fish = s.fish - amount
}

func (s *setup) updatePeriod() {
	s.period++
}

// GoFishing function start the simulatoin
func GoFishing(farmers map[int]interface{}, s *setup) {
	for i := 0; i < len(farmers); i++ {
		amountTaken := farmers[i].(farmer).amount
		s.updateFish(amountTaken)
		message := fmt.Sprintf("Farmer %d took %d fish, %d remaining", i, amountTaken, s.amountOfFish())
		logDetails(message)
	}
	s.createNewFish()
	s.updatePeriod()
}

func logDetails(message string) {
	display := color.New(color.FgGreen, color.Bold)
	display.Println(message)
}

// getInitialConditions returns the starting conditions of the system.
func (s setup) getInitialConditions() string {
	return fmt.Sprintf("There are %d farmers in the system with %d fish to begin with. \n%d new fish to be added at the end of the day.", s.farmers, s.fish, s.newFish)
}

// createFarmer method takes the number of farmers in the systen and creates an individual farmer which stored in a map
func (s setup) createFarmer() map[int]interface{} {
	var m = make(map[int]interface{})
	for i := 0; i < s.amountOfFarmers(); i++ {
		m[i] = farmer{true, 1}
	}
	display := color.New(color.FgWhite, color.Bold)
	display.Println("Farmers created!")
	return m
}

// this method creates a new fish for every two fish in the pond.
func (s *setup) createNewFish() {
	display := color.New(color.FgCyan, color.Bold)
	display.Println("Re-populating pond...")
	// check if there is less than 2 fish in the pond. If true, then an error id thrown and the program closes.
	if s.fish < 2 {
		log.Fatal("Can't repopulate the pond. \nNo more fish to reproduce :(. GAME OVER!")
	}

	newOnes := math.Floor(float64(s.fish / 2))
	s.fish = s.fish + int(newOnes)
	message := fmt.Sprintf("%d new fish added. Total is %d", int(newOnes), s.amountOfFish())
	logDetails(message)
}

// Welcome message to be displayed to the user
func Welcome() {
	fmt.Println("This is a simulatoin of the tragedy of the commons, a famous thought experiment which investigates what happens when individuals choose to go against what is good for the collective. ")
}
