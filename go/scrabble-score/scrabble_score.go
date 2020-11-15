// Scrabble score task of Exercism.io Go track.
package scrabble

import "strings"

// Compute scrabble score of the word.
func Score(word string) int {
	data := []struct {
		letters string
		score   int
	}{
		{"AEIOULNRST", 1},
		{"DG", 2},
		{"BCMP", 3},
		{"FHVWY", 4},
		{"K", 5},
		{"JX", 8},
		{"QZ", 10},
	}
	score := make(map[rune]int)
	for _, d := range data {
		for _, r := range d.letters {
			score[r] = d.score
		}
	}
	res := 0
	for _, r := range strings.ToUpper(word) {
		res += score[r]
	}
	return res
}
