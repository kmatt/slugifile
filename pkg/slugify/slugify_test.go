package slugify

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	var test = "La Quête d'Ewilan vol.1 : D'un monde à l'autre ·/_,:; (1), [Bottero, Pierre]`Author` @{1} <book> ?!//&"
	var expect = "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre._.1.bottero.pierre.author.1.book"

	var slug = Slugify(test)

	if slug != expect {
		t.Error("Expected 1, got ", slug)
	}
}
