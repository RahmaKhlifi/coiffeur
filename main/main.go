package main

import (
	"fmt"
	"log"
	"main/coiffeur"
)

func main() {
	sim := coiffeur.NewSimulator()

	sim.AddServer("Coiffeur1", 1.75)
	sim.AddServer("Coiffeur2", 1.5)
	sim.AddServer("Coiffeur3", 1)

	// Generate random arrival times based on Poisson distribution
	lambda := 1.0    // average arrival rate (clients per hour)
	duration := 10.0 // simulation duration (hours)

	// Use the function from poisson_test.go to generate arrival times
	arrivalTimes := coiffeur.GetArrivalTimesGivenServiceTime(lambda, duration)

	// Log generated arrival times for debugging
	log.Println("Generated arrival times:")
	for i, t := range arrivalTimes {
		log.Printf("Client %d: %.2f hour\n", i+1, t)
	}

	// Run the simulation with the generated arrival times
	sim.Run(arrivalTimes)
	fmt.Print("fin")
}
