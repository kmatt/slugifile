package listing

import (
	"testing"
)

func TestListing(t *testing.T) {
	var directory = "../../misc"
	var l = Scan(directory)

	if l.Count != 5 {
		t.Errorf("Expected 4 files, got %d", l.Count)
	}

	var lastFile = l.Files[l.Count-1]

	if lastFile.BasePath != "../../misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md" {
		t.Errorf("Expected ../../misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md, got %s", lastFile.BasePath)
	}

	if lastFile.Path != "/Users/ewilan/Workspace/go_tools/slugifier/misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md" {
		t.Errorf("Expected /Users/ewilan/Workspace/go_tools/slugifier/misc/listing/nested/La Quête d'Ewilan vol.1 D'un monde à l'autre.md, got %s", lastFile.Path)
	}

	if lastFile.Extension != ".md" {
		t.Errorf("Expected .md, got %s", lastFile.Extension)
	}

	if lastFile.FileName != "La Quête d'Ewilan vol.1 D'un monde à l'autre.md" {
		t.Errorf("Expected La Quête d'Ewilan vol.1 D'un monde à l'autre.md, got %s", lastFile.FileName)
	}
}
