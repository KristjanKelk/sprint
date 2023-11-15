package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	// Remove leading and trailing whitespaces, and convert to lowercase
	a := strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	b := strings.ToLower(strings.ReplaceAll(str2, " ", ""))
	// Convert strings to slices of characters
	aChars := strings.Split(a, "")
	bChars := strings.Split(b, "")

	// Sort the slices
	sort.Strings(aChars)
	sort.Strings(bChars)

	// Join the sorted slices back into strings
	sortedA := strings.Join(aChars, "")
	sortedB := strings.Join(bChars, "")

	// Compare the sorted strings
	return sortedA == sortedB
}
