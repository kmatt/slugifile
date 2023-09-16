package slugify

import (
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

type Replacement struct {
	To   string
	From *regexp.Regexp
}

// Slugify a string
func Slugify(input string) string {
	var seperator = "."
	var slug = unidecode.Unidecode(input)
	slug = strings.ToLower(slug)
	slug = strings.TrimSpace(slug)

	var sets = []Replacement{
		{To: "-", From: regexp.MustCompile(" - ")},
		{To: "_", From: regexp.MustCompile(" _ ")},
		{To: seperator, From: regexp.MustCompile(`[^-\w\s]`)}, // Remove all special characters except -
		{To: seperator, From: regexp.MustCompile(`[/^ /]`)},   // Replace all spaces with a dot
		{To: seperator, From: regexp.MustCompile("`")},
		{To: seperator, From: regexp.MustCompile(`[\s]+`)}, // Replace all spaces with a seperator
		{To: ".", From: regexp.MustCompile(`[.]+`)},        // Replace multiple dots with a single dot
		{To: "", From: regexp.MustCompile(`[.]$`)},         // Remove last dot
		{To: "", From: regexp.MustCompile(`^[.]`)},         // Remove first dot
	}

	for _, set := range sets {
		slug = set.From.ReplaceAllString(slug, set.To)
	}

	return slug
}
