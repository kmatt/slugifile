package file

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFile(t *testing.T) {
	var fileName = "La Quête d'Ewilan vol.1 D'un monde à l'autre"
	var prefix = "../../misc/file"
	var path = prefix + "/" + fileName + ".md"
	fmt.Println(path)

	var f = Scan(path)
	f.RenameAsSlug(true)

	equalsTo(f.BasePath, prefix, t)
	equalsTo(f.Extension, "md", t)
	equalsTo(f.FileName, "La Quête d'Ewilan vol.1 D'un monde à l'autre", t)
	equalsTo(strconv.FormatBool(f.IsExists), "true", t)
	equalsTo(f.SlugPath, prefix+"/la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre-1.md", t)
	equalsTo(f.SlugBasename, "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre-1.md", t)
	equalsTo(f.Path, prefix+"/La Quête d'Ewilan vol.1 D'un monde à l'autre.md", t)
	equalsTo(f.Slug, "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre-1", t)

	f.Rename(f.SlugPath, f.Path)
}

func equalsTo(a string, b string, t *testing.T) bool {
	var isEqual = a == b

	if !isEqual {
		t.Error("Expected " + a + ", got " + b)
	}

	return isEqual
}
