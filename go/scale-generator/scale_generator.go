// Scale generator task of Exercism.io Go track.
package scale

import "strings"

// Set of strings.
// found here: https://www.davidkaya.com/sets-in-golang/
type set struct {
	m map[string]struct{}
}

// Make set of strings and populate it with optional data.
func makeSet(items ...string) *set {
	s := &set{}
	s.m = make(map[string]struct{})
	for _, i := range items {
		s.put(i)
	}
	return s
}

// Put string into set.
func (s *set) put(item string) {
	s.m[item] = struct{}{}
}

// Test if set contains given string.
func (s *set) has(item string) bool {
	_, ok := s.m[item]
	return ok
}

// Given tonic and interval return a scale.
func Scale(tonic, interval string) (scale []string) {
	//sharps := makeSet("G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#")
	flats := makeSet("F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb")
	const (
		Sharps = 0
		Flats  = 1
	)
	kind := Sharps
	if flats.has(tonic) {
		kind = Flats
	}
	notes := [][]string{
		{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}, // sharps
		{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}} // flats
	if len(interval) == 0 {
		interval = "mmmmmmmmmmmm" // chromatic scale
	}
	return makeScale(tonic, interval[:len(interval)-1], notes[kind])
}

// Make a scale out of given tonic, interval string and line of possible notes.
func makeScale(tonic, interval string, notes []string) (scale []string) {
	i := find(tonic, notes)
	scale = append(scale, notes[i])
	step := map[rune]int{'m': 1, 'M': 2, 'A': 3}
	for _, r := range interval {
		i = (i + step[r]) % len(notes)
		scale = append(scale, notes[i])
	}
	return scale
}

// Return index of string 's' in slice 'notes'.
// Return -1 if not found.
// Compare strings case insensitive.
func find(s string, notes []string) int {
	for i := range notes {
		if strings.ToUpper(s) == strings.ToUpper(notes[i]) {
			return i
		}
	}
	///fmt.Printf("find(): %v not found in %v\n", s, notes)
	return -1
}
