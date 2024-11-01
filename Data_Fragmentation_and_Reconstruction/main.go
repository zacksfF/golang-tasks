//d-f-r stands for Data Fragmentation and Reconstruction Challenge 01
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// simpleHash generates a 30-character fixed-length hash from the input string.
func simpleHash(data string) string {
	hashValus := 0

	//Generate hash value by by iterating over each character
	for i := 0; i < len(data); i++ {
		hashValus += int(data[i]) * 31 % 100000000
	}

	//Covert HashValue to astring and adjust it to be exactly 30 characters long
	hashString := strconv.Itoa(hashValus)
	hashString = strings.Repeat(hashString, (30/len(hashString))+1)[:30]
	return hashString
}

// reconstruct_data check each fragment's integrity by comparing
// its stored hash with the recalculated hash. If any fragment
// fails verification, the  reconstruct_data function should handle
// it gracefully with an error message.
// reconstructData reassembles the original data from fragments, verifying integrity
func reconstructData(fragments map[int]map[string]string) (string, error) {
	// Create a slice to hold fragments in order and construct the data
	var reconstructedData string

	for i := 1; i <= len(fragments); i++ {
		fragment := fragments[i]
		data, hash := fragment["data"], fragment["hash"]

		// Verify the integrity of each fragment
		if simpleHash(data) != hash {
			return "", fmt.Errorf("Error: Data integrity verification failed")
		}

		// Append data fragment if hash matches
		reconstructedData += data
	}

	return reconstructedData, nil
}

func main() {
	// Define the fragments as a map of maps
	fragments := map[int]map[string]string{
		1: {"data": "Hello", "hash": simpleHash("Hello")},
		2: {"data": "Golang", "hash": simpleHash("Golang")},
		3: {"data": "!", "hash": simpleHash("!")},
	}

	// Attempt to reconstruct the data
	originalData, err := reconstructData(fragments)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reconstructed Data:", originalData)
	}
}
