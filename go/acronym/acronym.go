// This is 'acronym' task of Exercism.io go track.
package acronym

import "unicode"

// Abbreviate returns acronim made out of text passed to it.
func Abbreviate(s string) string {
	res := ""
	inWord := false
	for _, r := range s {
		if unicode.IsLetter(r) || r == '\'' {
			// the "|| r == '\''" part is needed to catch cases
			// like "Halley's Comet"
			if !inWord {
				inWord = true
				res += string(unicode.ToUpper(r))
			}
		} else {
			inWord = false
		}
	}
	return res
}
