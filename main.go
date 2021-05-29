// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that is a guessing game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This main function will be a better guessing game
func main() {
	// Defining variables
	var guess int
	rand.Seed(time.Now().UnixNano())
	var random int
	random = rand.Intn(9)

	fmt.Println("Better Guessing Game")
	fmt.Println()

	// Checks if guess == random
	for guess != random {
		// User Input
		fmt.Print("Guess(0-9): ")
		fmt.Scanln(&guess)
		fmt.Println()
		if guess < random {
			fmt.Println("Higher")
		} else if guess > random {
			fmt.Println("Lower")
		}
	}
	fmt.Println("You got it, number is", random)
}
