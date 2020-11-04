// Scale generator task of Exercism.io Go track.
package scale

import "strings"

// Given tonic and interval return a scale.
func Scale(tonic, interval string) []string {
	//sharps := []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
	flats := []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	const (
		Sharps = 0
		Flats  = 1
	)
	kind := Sharps
	if indexOf(tonic, flats, match) != -1 {
		kind = Flats
	}
	notes := [][]string{ // order does matter
		{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}, // sharps
		{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}} // flats
	if len(interval) == 0 {
		interval = "mmmmmmmmmmmm" // chromatic scale
	}
	return makeScale(tonic, interval[:len(interval)-1], notes[kind])
}

// Make a scale out of given tonic, interval string and line of possible notes.
func makeScale(tonic, interval string, notes []string) (scale []string) {
	i := indexOf(tonic, notes, matchCI)
	scale = append(scale, notes[i])
	step := map[rune]int{'m': 1, 'M': 2, 'A': 3}
	for _, r := range interval {
		i = (i + step[r]) % len(notes)
		scale = append(scale, notes[i])
	}
	return scale
}

// Function that returns true if strings a and b are the same.
type matchF func(a, b string) bool

// Case-sensitive strings comparison.
func match(a, b string) bool {
	return a == b
}

// Case-insensitive strings comparison.
func matchCI(a, b string) bool {
	return strings.ToUpper(a) == strings.ToUpper(b)
}

// Return index of string 's' in slice 'notes'.
// Return -1 if not found.
// Compare strings case insensitive.
func indexOf(s string, notes []string, match matchF) int {
	for i := range notes {
		if match(s, notes[i]) {
			return i
		}
	}
	///fmt.Printf("indexOf(): %v not found in %v\n", s, notes)
	return -1
}
