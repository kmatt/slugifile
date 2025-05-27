package renamer

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kmatt/slugifile/pkg/file"
	"github.com/kmatt/slugifile/pkg/listing"
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
					var levelStr = strconv.FormatInt(int64(l.Level), 10)
					log.Println("From " + "`" + f.FileName + "`")
					log.Println("To " + "`" + f.SlugBasename + "`" + " (" + "L." + levelStr + ")")
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
		log.Println("From " + "`" + f.FileName + "`")
		log.Println("To " + "`" + f.SlugBasename + "`")
		fmt.Println("")
	}

	return maxLevel
}

func Execute(filePath string, level int, lowercase bool) {
	var isDirectory = isDir(filePath)
	var i = level

	if isDirectory {
		for i > 0 {
			executeDirectory(filePath, i, lowercase)
			i--
		}

		handleDirectory(HandleDirectoryParams{
			FilePath: filePath,
			Callback: func(f file.File, l listing.ListingFile) {
				// Don't touch hidden files
				if !f.IsDir && !strings.HasPrefix(f.FileName, ".") && f.Path != filePath {
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
			if f.IsDir && !strings.HasPrefix(f.FileName, ".") && f.Path != filePath {
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
		log.Println("Error:", err)
		return false
	}

	return fileInfo.IsDir()
}
