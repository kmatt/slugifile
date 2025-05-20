package renamer

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kmatt/slugifier/pkg/file"
	"github.com/kmatt/slugifier/pkg/listing"
)

type Options struct {
	Verbose   bool
	Lowercase bool
}

func Preview(filePath string, options Options) int {
	var maxLevel = 0
	var isDirectory = isDir(filePath)

	if isDirectory {
		var list = handleDirectory(HandleDirectoryParams{
			FilePath: filePath,
			Callback: func(f file.File, l listing.ListingFile) {
				if options.Verbose {
					fmt.Println("From " + "`" + f.FileName + "`")
					var levelStr = strconv.FormatInt(int64(l.Level), 10)
					fmt.Println("To " + "`" + f.SlugBasename + "`" + " (" + "L." + levelStr + ")")
					fmt.Println("")
				}
			},
		}, options.Lowercase)

		for _, f := range list.Files {
			if f.IsDir {
				if f.Level > maxLevel {
					maxLevel = f.Level
				}
			}
		}
	} else {
		var f = file.Scan(filePath, options.Lowercase)
		fmt.Println("From " + "`" + f.FileName + "`")
		fmt.Println("To " + "`" + f.SlugBasename + "`")
		fmt.Println("")
	}

	return maxLevel
}

func Execute(filePath string, level int, lowercase bool) {
	var isDirectory = isDir(filePath)
	var i = level

	if isDirectory {
		for i > 0 {
			fmt.Println(i)
			executeDirectory(filePath, i, lowercase)
			i--
		}

		handleDirectory(HandleDirectoryParams{
			FilePath: filePath,
			Callback: func(f file.File, l listing.ListingFile) {
				// Don't touch hidden files
				if !f.IsDir && f.Path != filePath && !strings.HasPrefix(filePath, ".") {
					f.RenameAsSlug(true)
				}
			},
		}, lowercase)
	} else {
		var f = file.Scan(filePath, lowercase)
		f.RenameAsSlug(true)
	}
}

func executeDirectory(filePath string, level int, lowercase bool) {
	handleDirectory(HandleDirectoryParams{
		FilePath: filePath,
		Callback: func(f file.File, l listing.ListingFile) {
			if f.IsDir && f.FileName != "." && f.FileName != ".." && f.FileName != ".DS_Store" && f.FileName != ".git" && f.Path != filePath {
				f.RenameAsSlug(false)
			}
		},
		Level: level,
	}, lowercase)
}

type HandleDirectoryCallback func(file.File, listing.ListingFile)

type HandleDirectoryParams struct {
	FilePath string
	Callback HandleDirectoryCallback
	Level    int
}

// Handle if directory
func handleDirectory(params HandleDirectoryParams, lowercase bool) listing.Listing {
	var list = listing.Scan(params.FilePath)
	for _, f := range list.Files {
		var fs = file.Scan(f.Path, lowercase)

		if params.Level != 0 {
			if f.Level == params.Level {
				params.Callback(fs, f)
			}
		} else {
			params.Callback(fs, f)
		}
	}

	return list
}

// Check if the path is a directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return fileInfo.IsDir()
}
