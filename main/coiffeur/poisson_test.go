package coiffeur

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

func TestGetGivenServiceTime(t *testing.T) {
	rand.Seed(123456789)

	log.Printf("Start TestGetGivenServiceTime")
	arrivalTimes := GetArrivalTimesGivenServiceTime(1.5, 5.0)
	expected := []float64{.44, 1.68, 2.23, 2.41, 2.42, 3.76, 4.46, 4.97}

	assertEqualFloats(expected, arrivalTimes)
	log.Println("Arrival times of clients:")
	for i, t := range arrivalTimes {
		log.Printf("Client %d: %.2f hour\n", i+1, t)
	}
}

func TestGetGivenClientNumber(t *testing.T) {
	rand.Seed(123456789)

	log.Printf("Start TestGetGivenClientNumber")
	arrivalTimes := GetArrivalTimesGivenClientsNumber(0.5, 10)
	expected := []float64{1.32, 5.04, 6.68, 7.24, 7.27, 11.29, 13.39, 14.90, 17.94, 20.33}
	assertEqualFloats(expected, arrivalTimes)
	log.Println("Arrival times of clients:")
	for i, t := range arrivalTimes {
		log.Printf("Client %d: %.2f hour\n", i+1, t)
	}
}

func assertEqualFloats(expected, actual []float64) {
	if len(expected) != len(actual) {
		log.Fatalf("length of expected : %v != length of actual : %v\n", len(expected), len(actual))
	}

	for i, _ := range expected {
		if math.Abs(expected[i]-(actual[i])) > 0.01 {
			log.Fatalf("floats do not match: %v != %v\n",
				expected[i], actual[i])
		}
	}
}

/*func GetArrivalTimesGivenServiceTime(lambda float64, Time float64) []float64 {
	var arrivalTimes []float64
	time := 0.0

	for {
		interArrival := rand.ExpFloat64() / lambda
		time += interArrival
		if time > Time {
			break
		}
		arrivalTimes = append(arrivalTimes, time)
	}
	return arrivalTimes
}

func GetArrivalTimesGivenClientNumber(lambda float64, nbrclient int) []float64 {
	var arrivalTimes []float64
	time := 0.0

	for i := 0; i < nbrclient; i++ {
		interArrival := rand.ExpFloat64() / lambda
		time += interArrival
		arrivalTimes = append(arrivalTimes, time)
	}
	return arrivalTimes
}
*/
