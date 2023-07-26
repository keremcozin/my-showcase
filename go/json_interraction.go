// In this small program I wanted to show how Go can interract with JSON files.
// To do this we need to have a new folder "json" inside our root folder which
// holds the JSON data.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Worker struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	IsStudent bool   `json:"is_student"`
	Grades    []int  `json:"grades"`
	Address   struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
}

func main() {
	file, err := os.Open("./json/worker.json")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	var workerData Worker

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&workerData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Name:", workerData.Name)
	fmt.Println("Age:", workerData.Age)
	fmt.Println("Email:", workerData.Email)
	fmt.Println("Is Student:", workerData.IsStudent)
	fmt.Println("Grades:", workerData.Grades)
	fmt.Println("Address:")
	fmt.Println("  Street:", workerData.Address.Street)
	fmt.Println("  City:", workerData.Address.City)
	fmt.Println("  Zipcode:", workerData.Address.Zipcode)
}
