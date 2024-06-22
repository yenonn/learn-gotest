package main

import (
	"fmt"
)

func isPrime(n int) (bool, string) {
	// by definition, 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}
	// negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprintf("Negative numbers are not prime, %d is not a positive number", n)
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is prime!", n)
}

func intro() {
	fmt.Println("Prime number checker")
	prompt()
}

func prompt() {
	fmt.Print("Enter a number to check if it is prime: ")
}

func readUserInput(doneChan chan bool) {
	var n int
	for {
		fmt.Scanf("%d", &n)
		result, msg := isPrime(n)
		if result {
			fmt.Println(msg)
			doneChan <- result
			return
		} else {
			prompt()
		}
	}
}

func main() {
	doneChan := make(chan bool)
	defer close(doneChan)
	intro()
	go readUserInput(doneChan)
	<-doneChan
}
