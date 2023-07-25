// Simple Greeting Program in Go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func greetUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s! Thanks for checking my showcase.\n", name)
}

func main() {
	greetUser()
}
