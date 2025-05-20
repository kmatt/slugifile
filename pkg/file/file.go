package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/kmatt/slugifier/pkg/slugify"
)

type File struct {
	Path         string
	Extension    string
	FileName     string
	BasePath     string
	SlugPath     string
	Slug         string
	SlugBasename string
	IsExists     bool
	IsDir        bool
}

func Scan(path string, lowercase bool) File {
	var f = File{
		Path: path,
	}

	f.IsDir = isDir(f.Path)
	if !f.IsDir {
		f.Extension = f.setExtension()
	}
	f.FileName = f.setFileName()
	f.BasePath = f.setBasePath()
	f.IsExists = isFileExists(f.Path)
	f.Slug = slugify.Slugify(f.FileName, lowercase)

	if f.IsDir {
		f.SlugBasename = f.Slug
		f.SlugPath = f.BasePath + "/" + f.SlugBasename
	} else {
		f.SlugBasename = f.Slug + "." + f.Extension
		f.SlugPath = f.BasePath + "/" + f.SlugBasename
	}

	if !f.IsDir && f.SlugPath != f.Path {
		// var i = 1
		// Check if the new file path already exists
		// fmt.Println("File already exists:", f.SlugPath)
		// for isFileExists(f.SlugPath) {
		// f.Slug = f.Slug + "-" + fmt.Sprintf("%d", i)
		// f.SlugBasename = f.Slug + "." + f.Extension
		// f.SlugPath = f.BasePath + "/" + f.SlugBasename
		// i++
		// }
	}

	return f
}

func (f File) Rename(path string, newPath string) {
	err := os.Rename(path, newPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func (f File) RenameAsSlug(withError bool) {
	err := os.Rename(f.Path, f.SlugPath)

	if withError {
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func (f File) setExtension() string {
	var parts = strings.Split(f.Path, ".")
	var extension = parts[len(parts)-1]

	return extension
}

func (f File) setFileName() string {
	var parts = strings.Split(f.Path, "/")
	var fileName = parts[len(parts)-1]

	if !f.IsDir {
		// remove extension
		var parts2 = strings.Split(fileName, ".")
		fileName = strings.Join(parts2[:len(parts2)-1], ".")
	}

	return fileName
}

func (f File) setBasePath() string {
	var parts = strings.Split(f.Path, "/")
	var basePath = strings.Join(parts[:len(parts)-1], "/")

	return basePath
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			fmt.Println("Error:", err)
			return false
		}
	}

	return true
}

func isDir(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {

		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("Error:", err)
			return false
		}

		return fileInfo.IsDir()
	}

	return false
}
