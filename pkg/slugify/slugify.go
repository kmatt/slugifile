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
func Slugify(input string, lowercase bool) string {
	var separator = "-"
	var slug = unidecode.Unidecode(input)
	if lowercase {
		slug = strings.ToLower(slug)
	}
	slug = strings.TrimSpace(slug)

	var sets = []Replacement{
		{To: "", From: regexp.MustCompile("'")},
		{To: "-", From: regexp.MustCompile(" - ")},
		{To: "_", From: regexp.MustCompile(" _ ")},
		{To: separator, From: regexp.MustCompile(`[^-\w\s{}]+`)}, // Remove all special characters except -
		{To: separator, From: regexp.MustCompile(`[/^ /]`)},      // Replace all spaces with separator
		{To: separator, From: regexp.MustCompile("`")},
		{To: separator, From: regexp.MustCompile(`[\s]+`)}, // Replace all spaces with separator
		{To: ".", From: regexp.MustCompile(`[.]+`)},        // Replace multiple dots with separator
		{To: "", From: regexp.MustCompile(`[.]$`)},         // Remove last dot
		{To: "", From: regexp.MustCompile(`^[.]`)},         // Remove first dot
	}

	for _, set := range sets {
		slug = set.From.ReplaceAllString(slug, set.To)
	}

	return slug
}
