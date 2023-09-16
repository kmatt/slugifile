package slugify

import (
	"fmt"
	"testing"
)

func TestSlugify(t *testing.T) {
	type Testing struct {
		From string
		To   string
	}

	var sets = []Testing{
		{From: "La Quête d'Ewilan vol.1 : D'un monde à l'autre·/_,:; (1), [Bottero, Pierre]`Author` @{1} <book> ?!//&", To: "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre._.1.bottero.pierre.author.1.book"},
		{From: "00 - Préface", To: "00.preface"},
		{From: "Book - Author", To: "Book-Author"},
	}

	for _, set := range sets {
		var slug = Slugify(set.To)
		fmt.Println(set.To)

		if slug != set.To {
			t.Error("Expected " + set.To + " got " + slug)
		}
	}

}
