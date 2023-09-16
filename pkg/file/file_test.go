package file

import (
	"log"
	"os"
	"strconv"
	"testing"
)

func TestFile(t *testing.T) {
	var fileName = "La Quête d'Ewilan vol.1 D'un monde à l'autre"
	var prefix = "../.."
	var path = prefix + "/" + fileName + ".txt"

	_, e := os.Create(path)
	if e != nil {
		log.Fatal(e)
	}

	var f = Scan(path)
	f.RenameAsSlug(true)

	equalsTo(f.BasePath, prefix, t)
	equalsTo(f.Extension, "txt", t)
	equalsTo(f.FileName, "La Quête d'Ewilan vol.1 D'un monde à l'autre", t)
	equalsTo(strconv.FormatBool(f.IsExists), "true", t)
	equalsTo(f.SlugPath, prefix+"/la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre.txt", t)
	equalsTo(f.SlugBasename, "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre.txt", t)
	equalsTo(f.Path, prefix+"/La Quête d'Ewilan vol.1 D'un monde à l'autre.txt", t)
	equalsTo(f.Slug, "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre", t)

	f.Rename(f.SlugPath, f.Path)

	er := os.Remove(f.Path)
	if er != nil {
		log.Fatal(er)
	}
}

func equalsTo(a string, b string, t *testing.T) bool {
	var isEqual = a == b

	if !isEqual {
		t.Error("Expected " + a + ", got " + b)
	}

	return isEqual
}
