package renamer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ewilan-riviere/slugifier-cli/pkg/file"
	"github.com/ewilan-riviere/slugifier-cli/pkg/listing"
)

func Preview(filePath string, verbose bool) int {
	var maxLevel = 0
	var isDirectory = isDir(filePath)

	if isDirectory {
		var list = handleDirectory(HandleDirectoryParams{
			FilePath: filePath,
			Callback: func(f file.File, l listing.ListingFile) {
				if verbose {
					fmt.Println("From " + "`" + f.FileName + "`")
					var levelStr = strconv.FormatInt(int64(l.Level), 10)
					fmt.Println("To " + "`" + f.SlugBasename + "`" + " (" + "L." + levelStr + ")")
					fmt.Println("")
				}
			},
		})

		for _, f := range list.Files {
			if f.IsDir {
				if f.Level > maxLevel {
					maxLevel = f.Level
				}
			}
		}
	} else {
		var f = file.Scan(filePath)
		fmt.Println("From " + "`" + f.FileName + "`")
		fmt.Println("To " + "`" + f.SlugBasename + "`")
		fmt.Println("")
	}

	return maxLevel
}

func Execute(filePath string, level int) {
	var isDirectory = isDir(filePath)
	var i = level

	if isDirectory {
		for i > 0 {
			fmt.Println(i)
			executeDirectory(filePath, i)
			i--
		}

		handleDirectory(HandleDirectoryParams{
			FilePath: filePath,
			Callback: func(f file.File, l listing.ListingFile) {
				if !f.IsDir && f.FileName != "." && f.FileName != ".." && f.FileName != ".DS_Store" && f.FileName != ".git" && f.Path != filePath {
					f.RenameAsSlug(true)
				}
			},
		})
	} else {
		var f = file.Scan(filePath)
		f.RenameAsSlug(true)
	}
}

func executeDirectory(filePath string, level int) {
	handleDirectory(HandleDirectoryParams{
		FilePath: filePath,
		Callback: func(f file.File, l listing.ListingFile) {
			if f.IsDir && f.FileName != "." && f.FileName != ".." && f.FileName != ".DS_Store" && f.FileName != ".git" && f.Path != filePath {
				f.RenameAsSlug(false)
			}
		},
		Level: level,
	})
}

type HandleDirectoryCallback func(file.File, listing.ListingFile)

type HandleDirectoryParams struct {
	FilePath string
	Callback HandleDirectoryCallback
	Level    int
}

// Handle if directory
func handleDirectory(params HandleDirectoryParams) listing.Listing {
	var list = listing.Scan(params.FilePath)
	for _, f := range list.Files {
		var fs = file.Scan(f.Path)

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
