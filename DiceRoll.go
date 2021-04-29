// This program will generate a random number and print to console.
// Created By: Marlon Poddalgoda
// Created on: 2021-04-28

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RollDie(maxInput int) {
	// this function generates and prints a random number

	// generates random number
	rand1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(rand1)

	// output random number to console
	fmt.Println("")
	fmt.Println(random.Intn(maxInput))
}

func main() {
	// this function takes in user input

	fmt.Println("This program generates a random number from a certain range.")
	fmt.Println("Enter the range (0 to ?).")
	fmt.Println("")

	// variables
	var maxInput string
	var maxValue int

	fmt.Print("How high do you want the number range to be (0 to ?): ")

	// check user input
	for {
		// user input
		_, err := fmt.Scan(&maxInput)
		maxValue, err = strconv.Atoi(maxInput)

		// if error, then restart loop
		switch {
		case err != nil:
			// output
			fmt.Println("That's not a  number, try again!")
			fmt.Print("How high do you want the number range to be (0 to ?): ")
			return
		case maxValue <= 0:
			// output
			fmt.Println("Please enter a positive number, try again!")
			fmt.Print("How high do you want the number range to be (0 to ?): ")
		case maxValue > 0:
			// calls procedure
			RollDie(maxValue)
			break
		default:
		    // error catch
		    fmt.Println("Sorry, please try again!")
	    }
	    // leave for loop
	    break
	}
	// program closes
	fmt.Println("")
	fmt.Println("Done.")
}
