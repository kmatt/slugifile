package main

import (
	"flag"
	"fmt"
	"os"
	"slugifier/pkg/renamer"
	"slugifier/pkg/utils"
	"strings"
)

func main() {
	var showHelp bool
	var verbose = false
	flag.BoolVar(&showHelp, "h", false, "Display help message")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Welcome to slugifier-cli!")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintf(os.Stderr, "Example: `slugifier path/to/file` or `slugifier path/to/directory`")
		// fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] MESSAGE\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "")
	}

	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	var filePath = getPath()

	if filePath == "" {
		fmt.Println("No file path provided.")
		return
	}

	var maxLevel = renamer.Preview(filePath, verbose)
	if verbose {
		var accept = utils.Confirm()

		if !accept {
			return
		}
	}

	renamer.Execute(filePath, maxLevel)
	fmt.Println("Done!")
}

// Get the file path from the command-line arguments
func getPath() string {
	// Check if at least one command-line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: slugifier <file-path>")
		return ""
	}

	// The first argument (os.Args[0]) is the program name
	filePath := os.Args[1]

	if strings.HasPrefix(filePath, "-") {
		filePath = os.Args[2]
	}

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File '%s' does not exist.\n", filePath)
		} else {
			fmt.Println("Error:", err)
		}
		return ""
	}

	return filePath
}
