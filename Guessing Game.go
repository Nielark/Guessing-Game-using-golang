package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var guessInput int
	life := 10
	seed := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(seed)

	randNumGen := randNum.Intn(51-1) + 1 // Random number generator range from 1 to 50.

	fmt.Println("G U E S S I N G  G A M E")
	fmt.Println()

	for life != 0 && guessInput != randNumGen {
		fmt.Println("Life ---> " + strconv.Itoa(life)) // Shows the remaining life.
		fmt.Print("Enter your guess number(from 1 - 50): ")
		fmt.Scanln(&guessInput)
		fmt.Print()

		if guessInput > 0 && guessInput < 51 {
			if guessInput == randNumGen {
				fmt.Println("Your guess is right.")
			} else if guessInput < randNumGen {
				life--
				fmt.Println("Your guess is wrong, make your guess higher.")
			} else {
				life--
				fmt.Println("Your guess is wrong, make your guess lower.")
			}
		} else {
			fmt.Println("Invalid Input.")
		}
		fmt.Println()
	}

	if life == 0 {
		fmt.Print("You have no remaining life, better luck next time.")
	}
}
