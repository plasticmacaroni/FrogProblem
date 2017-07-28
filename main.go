package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//We need to try to make sure our numbers are as random as possible
	rand.Seed(time.Now().UTC().UnixNano())

	totalFemaleGroups := 0
	totalIterations := float64(10000000)

	//Let's create a group of two frogs, where at least one frog is male
	for x := 0; x < int(totalIterations); x++ {
		frogSet := createFrogs()
		//If either of our two frogs are female, we need to increase the total female counter by one
		if frogSet[0] == 0 || frogSet[1] == 0 {
			totalFemaleGroups++
		}
	}
	fmt.Println(totalFemaleGroups, "of the "+fmt.Sprint(totalIterations)+" pairs included a female.")
	fmt.Println("That's", fmt.Sprint(float64(totalFemaleGroups)/totalIterations*100)+"%", "of the frog sets that were female!")
}

func createFrogs() []float64 {
	frogSet := []float64{0, 0}

	//We're going to say 0 is a female, and 1 is a male
	for x := 0; x < 2; x++ {
		frogSet[x] = math.Floor(rand.Float64() + .5)
	}
	//We need to throw out scenarios where we only ended up with females
	if frogSet[0] == 0 && frogSet[1] == 0 {
		frogSet = createFrogs()
	}
	return frogSet
}
