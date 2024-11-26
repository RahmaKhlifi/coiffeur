package coiffeur

import (
	//    "fmt"
	"math/rand"
)

// Returns a slice of arrival times given the Poisson process rate
// (lambda) and serviceTime, all arrival times should be less than serviceTime
// The interArrival time is given by the  following expression :
// rand.ExpFloat64() / lambda

func GetArrivalTimesGivenServiceTime(lambda float64, serviceTime float64) []float64 {
	arrivalTimes := make([]float64, 0)
	var currentTime float64 = 0.0

	for currentTime < serviceTime {
		// Generate the inter-arrival time
		interArrival := rand.ExpFloat64() / lambda
		currentTime += interArrival
		arrivalTimes = append(arrivalTimes, currentTime)
	}
	if arrivalTimes[len(arrivalTimes)-1] > serviceTime {
		return arrivalTimes[:len(arrivalTimes)-1]
	}
	return arrivalTimes
}

// Returns a slice of arrival times given the Poisson process rate
// (lambda) and the number of clients,
// The interArrival time is given by the  following expression :
// rand.ExpFloat64() / lambda
func GetArrivalTimesGivenClientsNumber(lambda float64, numClients int) []float64 {
	arrivalTimes := make([]float64, numClients)
	var currentTime float64 = 0.0
	for i := 0; i < numClients; i++ {
		// Generate the inter-arrival time
		interArrival := rand.ExpFloat64() / lambda
		currentTime += interArrival
		arrivalTimes[i] = currentTime
	}
	return arrivalTimes
}
