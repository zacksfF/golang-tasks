package main

import (
	"fmt"
	"math"
	"sort"
)

// CalculateRisk function to calculate risk for a given number of fragments in a data center
func CalculateRisk(risk, fragment int) int {
	return int(math.Pow(float64(risk), float64(fragment)))
}

// canDistribute its a functionm to chck if we can distibuted fragment to keep max risk under a given threshold
func canDistribute(dataCentres []int, fragment int, maxRisk int) bool {
	remainingFragment := fragment

	for _, risk := range dataCentres {
		allocated := 0

		for CalculateRisk(risk, allocated+1) <= maxRisk {
			allocated++
		}

		//Dedut allocated frafment for this data center
		remainingFragment -= allocated

		if remainingFragment <= 0 {
			return true
		}
	}
	return remainingFragment <= 0
}

// distributeFragments is a Main function to find the minimized maximum risk
func distributeFragments(dataCenters []int, fragments int) int {
	// Sort data centers to allow for optimal fragment allocation
	sort.Ints(dataCenters)

	// Define binary search bounds
	left := 0
	right := int(math.Pow(float64(dataCenters[len(dataCenters)-1]), float64(fragments)))

	// Binary search for the minimum achievable maximum risk
	for left < right {
		mid := left + (right-left)/2

		// Check if it's possible to distribute within this mid value as max risk
		if canDistribute(dataCenters, fragments, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	//usage
	dataCenters := []int{10, 20, 30}
	fragment := 5
	minRisk := distributeFragments(dataCenters, fragment)
	fmt.Printf("Minimized maximum risk: %d\n", minRisk)
}
