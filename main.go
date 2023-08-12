package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNumber()
	// fmt.Println(secret)
	for matching := false; !matching; {
		guess := getUserInput()
		fmt.Println(secret, guess)

		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Guess is too small")
		return false
	} else {
		fmt.Println("How did you guess that? You're right!")
		return true
	}
}

func getUserInput() int {
	fmt.Print("Enter your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Please enter valid input!")
	} else {
		fmt.Println("Your guess", input)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
