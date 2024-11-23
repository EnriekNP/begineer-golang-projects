package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(100) + 1

		var counter int
		for {
			if counter >= 3 {
				fmt.Printf("You lose! The correct number was %d\n", randomNumber)
				break
			}
			counter++

			fmt.Println("Guess The Number!!! (1-100)")
			var playerNumber int
			_, err := fmt.Scanln(&playerNumber)

			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}

			if playerNumber > 100 || playerNumber < 1 {
				fmt.Println("Invalid guess. Please enter a number between 1 and 100.")
				continue
			}
			if playerNumber > randomNumber {
				fmt.Println("Too high.")
			} else if playerNumber == randomNumber {
				fmt.Println("You guessed it!")
				break
			} else {
				fmt.Println("Too low.")
			}
		}
		fmt.Println("Do you want to play again? (yes/no)")
		var playAgain string
		fmt.Scanln(&playAgain)

		if strings.ToLower(playAgain) != "yes" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

	}
}
