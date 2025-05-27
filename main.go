// Package notifier is a small tool to slugify files and directories, recursively.
//
// Examples/readme can be found on the GitHub page at https://github.com/ewilan-riviere/slugifile
//
// If you want to use it as CLI, you can install it with:
//
//	go install github.com/ewilan-riviere/slugifile
//
// Then you can use it like this:
//
//	slugifile path/to/dir
//	slugifile path/to/file
//
// You can use `-l` flag to lowercase the slugified files and directories.
// You can use `-v` flag to enable verbose mode.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kmatt/slugifile/pkg/renamer"
	"github.com/kmatt/slugifile/pkg/slugify"
	"github.com/kmatt/slugifile/pkg/utils"
)

// Get the file path from the command-line arguments
func getPath() string {
	// Check if at least one command-line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: slugifile <file-path>")
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
			log.Println("Error:", err)
		}
		return ""
	}

	return filePath
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var showHelp = false
	var verbose = false
	var lowercase = false
	var text = ""

	flag.BoolVar(&showHelp, "h", false, "Display help message")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.BoolVar(&lowercase, "l", false, "Lowercase mode")
	flag.StringVar(&text, "t", "", "Slugify text and exit")

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

	if text != "" {
		fmt.Println(slugify.Slugify(text, lowercase))
		return
	}

	var filePath = getPath()

	if filePath == "" {
		log.Println("No file path provided")
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
	fmt.Println("Done")
}
