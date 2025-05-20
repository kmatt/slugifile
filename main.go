// Package notifier is a small tool to slugify files and directories, recursively.
//
// Examples/readme can be found on the GitHub page at https://github.com/ewilan-riviere/slugifier
//
// If you want to use it as CLI, you can install it with:
//
//	go install github.com/ewilan-riviere/slugifier
//
// Then you can use it like this:
//
//	slugifier path/to/dir
//	slugifier path/to/file
//
// You can use `-l` flag to lowercase the slugified files and directories.
// You can use `-v` flag to enable verbose mode.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ewilan-riviere/slugifier/pkg/renamer"
	"github.com/ewilan-riviere/slugifier/pkg/utils"
)


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

func main() {
	var showHelp = false
	var verbose = false
	var lowercase = false

	flag.BoolVar(&showHelp, "h", false, "Display help message")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.BoolVar(&lowercase, "l", false, "Lowercase mode")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] path/to/file-or-directory\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
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

	var maxLevel = renamer.Preview(filePath, renamer.Options{
		Verbose:   verbose,
		Lowercase: lowercase,
	})
	if verbose {
		var accept = utils.Confirm()

		if !accept {
			return
		}
	}

	renamer.Execute(filePath, maxLevel, lowercase)
	fmt.Println("Done!")
}
