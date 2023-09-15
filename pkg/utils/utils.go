package utils

import (
	"fmt"
	"os"
	"strings"
)

// confirm function asks for user input
func Confirm() bool {
	var input string

	fmt.Println("")
	fmt.Printf("Do you want to continue with this operation? [Y|n]: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		// panic(err)
		input = "y"
	}
	input = strings.ToLower(input)

	if input == "y" || input == "yes" {
		return true
	}

	if input == "n" || input == "no" {
		os.Exit(1)
		return false
	}

	return false
}
