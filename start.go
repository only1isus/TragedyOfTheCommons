package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type Setup struct {
	farmers int // the amont of farmers in the system
	fish    int // the amount of fish in the pond to begin with
	newFish int // the amount of new fishes to be added to the system at the end of a day.
	period  int // tracks the of 'days'
}

type Farmer struct {
	honest bool // tells if the farmer is being honest (choosing to abide by the rules) a farmer is 'honest' if he only takes one fish
	amount int  // the amount of fish the farmer takes from the pond.
}

func main() {
	Welcome()
	Start()
}

func Start() {
	for {
		systemSetup := Setup{4, 20, 2, 0} // this is where the starting conditions are set. farmers, fish, amount taken
		farmers := RandomizeFarmers(5, &systemSetup)
		systemSetup.details()
		GoFishing(farmers, &systemSetup)
		time.Sleep(2 * time.Second)
	}
}

// GoFishing function start the simulatoin
func GoFishing(farmers map[int]Farmer, s *Setup) {
	for i := 0; i < len(farmers); i++ {
		amountTaken := farmers[i].amount
		s.updateFish(amountTaken)
		message := fmt.Sprintf("Farmer %d took %d fish, %d remaining", i, amountTaken, s.amountOfFish())
		logDetails(message)
	}
	// quick and dirty, add new fish every 2 days instead of 1
	fmt.Println(s.period % 2)
	if s.period%2 == 0 {
		s.createNewFish()
	}
	s.updatePeriod()
}

func logDetails(message string) {
	display := color.New(color.FgGreen, color.Bold)
	display.Println(message)
}

/*
	RandomizeFarmers function creates a random farmer and assign randon atributes to that farmer.
	The result is returned as a map. The map returned can be empty.
	If a farmer is honest then he takes only one fisn. If he is not however then he would
	take two or more fish.
*/
func RandomizeFarmers(max int, s *Setup) map[int]Farmer {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	numberOfFarmers := 3 + rand.Intn(max-3)
	s.farmers = numberOfFarmers
	m := make(map[int]Farmer)

	for i := 0; i < numberOfFarmers; i++ {
		rand.Seed(time.Now().UnixNano())

		x := rand.Intn(2)
		if x == 1 {
			m[i] = Farmer{true, 1}
		} else {
			amount := 2 + rand.Intn(4-2)
			m[i] = Farmer{false, amount}
		}
	}
	return m
}

// details method tells the initial condition of the experiment
func (s *Setup) details() {
	display := color.New(color.FgCyan, color.Bold)
	display.Printf("%d farmers, %d fish available, %d days\n", s.amountOfFarmers(), s.amountOfFish(), s.period)
	// return fmt.Sprintf("Honest --> %t | Taking --> %d fish", f.honest, f.amount)
}

// this method returns the amount of farmers in the system
func (s *Setup) amountOfFarmers() int {
	return s.farmers
}

// this method returns the amount of fish in the system
func (s *Setup) amountOfFish() int {
	return s.fish
}

// this method updates the amount of fish in the pond
func (s *Setup) updateFish(amount int) {
	s.fish = s.fish - amount
}

func (s *Setup) updatePeriod() {
	s.period++
}

// getInitialConditions returns the starting conditions of the system.
func (s Setup) getInitialConditions() string {
	return fmt.Sprintf("There are %d farmers in the system with %d fish to begin with. \n%d new fish to be added at the end of the day.", s.farmers, s.fish, s.newFish)
}

// createFarmer method takes the number of farmers in the systen and creates an individual farmer which stored in a map
func (s Setup) createFarmer() map[int]interface{} {

	// create farmers with random fish count
	var m = make(map[int]interface{})
	for i := 0; i < s.amountOfFarmers(); i++ {
		m[i] = Farmer{true, 1}
	}
	display := color.New(color.FgWhite, color.Bold)
	display.Println("Farmers created!")
	return m
}

// this method creates a new fish for every two fish in the pond.
func (s *Setup) createNewFish() {
	display := color.New(color.FgCyan, color.Bold)
	display.Println("Re-populating pond...")
	// check if there is less than 2 fish in the pond. If true, then an error id thrown and the program closes.
	if s.fish < 2 {
		display := color.New(color.FgRed, color.Bold)
		display.Println("Can't repopulate the pond. \nNo more fish to reproduce :(. GAME OVER!")
		// log.Fatal("Can't repopulate the pond. \nNo more fish to reproduce :(. GAME OVER!")
		os.Exit(1)
	}

	newOnes := int(math.Floor(float64(s.fish/2))) * s.newFish
	s.fish = s.fish + int(newOnes)
	message := fmt.Sprintf("%d new fish added. Total is %d", int(newOnes), s.amountOfFish())
	logDetails(message)
}

// Welcome message to be displayed to the user
func Welcome() {
	fmt.Println("This is a simulatoin of the tragedy of the commons, a famous thought experiment which investigates what happens when individuals choose to go against what is good for the collective. ")
}
