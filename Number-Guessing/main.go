package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var user_input int

	attempts := 0
	max_attempts := 7
	won := false

	secret_number := rand.Intn(101)

	fmt.Println("I Picked a number between 0 - 100")
	fmt.Println("You have 7. attempts, Good luck.")

	for attempts < max_attempts {

		fmt.Println("Please enter your guess :")
		fmt.Scan(&user_input)

		attempts++

		if user_input < secret_number {

			fmt.Println("Too low, Try again.")

		} else if user_input > secret_number {

			fmt.Println("Too high, Try again.")

		} else {

			fmt.Println("Congratulations you won the game.")
			fmt.Println("You have taken:", attempts, "attempts")

			won = true
			break
		}
	}

	if !won {
		fmt.Println("You lost. Try again")
		fmt.Println("The secret number was:", secret_number)
	}
}