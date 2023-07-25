// A Go Program to calculate factorials concurrently with user given input.
// big.Int is used in this program to allow user to enter bigger numbers. 

package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64, result chan<- *big.Int) {
	fact := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	result <- fact
}

func main() {
	fmt.Print("Enter a number to calculate its factorial: ")
	var input string
	fmt.Scanln(&input)

	num, success := new(big.Int).SetString(input, 10)
	if !success {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	resultChan := make(chan *big.Int)

	go factorial(num.Int64(), resultChan)

	result := <-resultChan
	fmt.Printf("Factorial of %s: %s\n", num.String(), result.String())
}
