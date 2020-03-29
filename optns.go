package main

import (
	"fmt"
)

func showOpts(options []string) {
	fmt.Println("Please select an option:")
	for _, option := range options {
		fmt.Println(option)
	}
	var choice string
	fmt.Scan(&choice)
	displayResponse(choice)
}

func displayResponse(choice string) {
	switch choice {
	case "1":
		fmt.Printf("First the worst: You selected %v \n", choice)
	case "2":
		fmt.Printf("Second the Best: You selected %v \n", choice)
	case "3":
		fmt.Printf("Third the one with the hairy chest: You selected %v \n", choice)
	default:
		fmt.Printf("Sorry, %v is not a valid option \n", choice)
	}
}
