package main

import (
	"fmt"
	"testing"
)

func TestGoFishing(t *testing.T) {
	want := 12 // 12 expected to be in the pond at the end of the day
	t.Run(("Setup"), func(t *testing.T) {
		init := Setup{4, 12, 1, 0}                  // initialize the starting getInitialConditions
		init.updateFish(init.amountOfFarmers() * 1) // update the amount of fish taken form the pond
		init.createNewFish()
		// fmt.Printf("Got %d fish after fishing. Expected %d\n", init.amountOfFish(), want)
		if init.amountOfFish() != want {
			t.Error("The ampunt of fish at the end of the day doesn't match")
		}
	})
}

func TestStart(t *testing.T) {
	want := 16 // 12 expected to be in the pond at the end of the day
	t.Run(("Setup"), func(t *testing.T) {
		init := Setup{4, 12, 2, 0}                  // initialize the starting getInitialConditions
		init.updateFish(init.amountOfFarmers() * 1) // update the amount of fish taken form the pond
		init.createNewFish()
		// fmt.Printf("Got %d fish after fishing. Expected %d\n", init.amountOfFish(), want)
		if init.amountOfFish() != want {
			fmt.Printf("Got %d fish after fishing. Expected %d\n", init.amountOfFish(), want)
			t.Error("The amount of fish at the end of the day doesn't match.")
		}
	})
}
