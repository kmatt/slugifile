package slugify

import (
	"regexp"
	"strings"
)

type Replacement struct {
	To   string
	From *regexp.Regexp
}

// Slugify a string
func Slugify(input string) string {
	var slug = strings.ToLower(input)

	var sets = []Replacement{
		{To: "a", From: regexp.MustCompile(`[àáâãäåæāăąạảấầẩẫậắằẳẵặ]`)},
		{To: "c", From: regexp.MustCompile(`[çćĉč]`)},
		{To: "d", From: regexp.MustCompile(`[ðďđþ]`)},
		{To: "e", From: regexp.MustCompile(`[èéêëēĕėęěẹẻẽếềểễệ]`)},
		{To: "g", From: regexp.MustCompile(`[ĝğģǵ]`)},
		{To: "h", From: regexp.MustCompile(`[ĥḧ]`)},
		{To: "i", From: regexp.MustCompile(`[ìíîïĩīįi̇ỉị]`)},
		{To: "j", From: regexp.MustCompile(`[ĵ]`)},
		{To: "i", From: regexp.MustCompile(`'[ĳ]`)},
		{To: "k", From: regexp.MustCompile(`[ķ]`)},
		{To: "l", From: regexp.MustCompile(`[ĺļľł]`)},
		{To: "m", From: regexp.MustCompile(`[ḿ]`)},
		{To: "n", From: regexp.MustCompile(`[ñńņň]`)},
		{To: "o", From: regexp.MustCompile(`[òóôõöøōŏőọỏốồổỗộớờởỡợǫǭơ]`)},
		{To: "o", From: regexp.MustCompile(`'[œ]`)},
		{To: "p", From: regexp.MustCompile(`[ṕ]`)},
		{To: "r", From: regexp.MustCompile(`[ŕŗř]`)},
		{To: "s", From: regexp.MustCompile(`[ßśŝşš]`)},
		{To: "t", From: regexp.MustCompile(`[ţť]`)},
		{To: "u", From: regexp.MustCompile(`[ùúûüũūŭůűųụủứừửữựư]`)},
		{To: "w", From: regexp.MustCompile(`[ẃŵẁẅ]`)},
		{To: "x", From: regexp.MustCompile(`[ẍ]`)},
		{To: "y", From: regexp.MustCompile(`[ýŷÿỳỵỷỹ]`)},
		{To: "z", From: regexp.MustCompile(`[źżž]`)},
		{To: ".", From: regexp.MustCompile(`[^\w\s]`)}, // Replace all special chars with a dot
		{To: ".", From: regexp.MustCompile(`[/^ /]`)},  // Replace all spaces with a dot
		// {To: seperator, From: regexp.MustCompile(`[\s]+`)}, // Replace all spaces with a seperator
		{To: ".", From: regexp.MustCompile(`[.]+`)}, // Replace multiple dots with a single dot
		{To: "", From: regexp.MustCompile(`[.]$`)},  // Remove last dot
		{To: "", From: regexp.MustCompile(`^[.]`)},  // Remove first dot
	}

	for _, set := range sets {
		slug = set.From.ReplaceAllString(slug, set.To)
	}

	return slug
}
