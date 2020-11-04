// Exercism Go Track cource, Hamming distance exercise.
package hamming

import "fmt"

// Distance returns Hamming distance between a and b strings (i.e. number of
// mismatches between strings.
// In case if strings are of different length non-nil error is returned,
// otherwise error is nil.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strands lengths not the same")
	}
	n := 0
	for i, _ := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n, nil
}
