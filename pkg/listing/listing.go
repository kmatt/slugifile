package listing

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ListingFile struct {
	BasePath  string
	Path      string
	Extension string
	FileName  string
	IsDir     bool
	Level     int
}

type Listing struct {
	Directory string
	Files     []ListingFile
	Count     int
}

func Scan(directory string) Listing {
	var l = Listing{
		Directory: directory,
		Files:     []ListingFile{},
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var lFile = l.toFile(path, info.IsDir())
		if lFile.FileName != "." && lFile.FileName != ".." && lFile.FileName != ".DS_Store" && lFile.FileName != ".git" && lFile.Path != directory {
			l.Files = append(l.Files, lFile)
		}

		if lFile.FileName == ".DS_Store" {
			e := os.Remove(lFile.Path)
			if e != nil {
				log.Fatal(e)
			}
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	l.Count = len(l.Files)

	return l
}

func (l Listing) toFile(path string, isDir bool) ListingFile {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}

	var f = ListingFile{
		BasePath:  path,                        // `../../misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md`
		Path:      absolutePath,                // `/Users/ewilan/Workspace/go_tools/slugifile/misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md`
		Extension: filepath.Ext(absolutePath),  // `.md`
		FileName:  filepath.Base(absolutePath), // `La Quête d'Ewilan vol.1 D'un monde à l'autre.md`
		IsDir:     isDir,
		Level:     len(strings.Split(absolutePath, "/")) - len(strings.Split(l.Directory, "/")),
	}

	return f
}
