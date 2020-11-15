// Isogram task of Exercism.io Go track.
package isogram

import "unicode"

// Retrun true if word is isogram (does not contain duplicating letters).
func IsIsogram(word string) bool {
	count := make(map[rune]bool)
	for _, r := range word {
		r = unicode.ToUpper(r)
		if unicode.IsLetter(r) && count[r] {
			return false
		}
		count[r] = true
	}
	return true
}
