package slugify

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	type Testing struct {
		From string
		To   string
	}

	var sets = []Testing{
		{From: "La Quête d'Ewilan vol.1 : D'un monde à l'autre-·/_,:; (1), [Bottero, Pierre]`Author` @{1} <book> ?!//&", To: "la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre-._.1.bottero.pierre.author.1.book"},
		{From: "00 - Préface", To: "00-preface"},
		{From: "00 _ Préface", To: "00_preface"},
		{From: "Book - Author", To: "book-author"},
		{From: "Book-Author", To: "book-author"},
		{From: "Book..Author", To: "book.author"},
		{From: "@!&^*()[]{}#%$£\":;.,<>|’«»“”‘’‹›''\\", To: "ps"},
		{From: `""`, To: ""},
		{From: "Здравствуйте", To: "zdravstvuite"},
		{From: "Γεια σας", To: "geia.sas"},
		{From: "こんにちは", To: "konnichiha"},
		{From: "안녕하세요", To: "annyeonghaseyo"},
		{From: "你好", To: "ni.hao"},
		{From: "Olá", To: "ola"},
		{From: "Здраво", To: "zdravo"},
		{From: "Здравейте", To: "zdraveite"},
		{From: "Привет", To: "privet"},
		{From: "Ahoj", To: "ahoj"},
		{From: "Halo", To: "halo"},
		{From: "Salam", To: "salam"},
		{From: "Ciao", To: "ciao"},
		{From: "السلام عليكم", To: "lslm.lykm"},
		{From: "Góðan daginn", To: "godan.daginn"},
	}

	for _, set := range sets {
		var slug = Slugify(set.From)

		if slug != set.To {
			t.Error("Expected " + set.To + " got " + slug)
		}
	}
}
