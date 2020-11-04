// This is Proverb exercise of Exercism.io Go track.
package proverb

import "fmt"

// Generate proverb out of list of words.
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) == 0 {
		return
	}
	return append(doRhyme(proverb, rhyme), fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}

// Build proverb recursively (except for the last line).
func doRhyme(proverb []string, words []string) []string {
	if len(words) > 1 {
		return doRhyme(append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", words[0], words[1])), words[1:])
	}
	return proverb
}
