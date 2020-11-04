// Raindrops exercise of Exercism.io Go track.
package raindrops

// iota converts n into string (base 10).
// (For the sake of exercise and to get rid of external dependencies.)
func itoa(n int) string {
	r, q := n%10, n/10
	d := string(rune('0' + r))
	if q == 0 {
		return d
	}
	return itoa(q) + d
}

// Convert turns n into "raindrop" string.
func Convert(n int) string {
	if n == 0 {
		return itoa(n)
	}

	factors := []struct {
		factor int
		name   string
	}{ // order does matter
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"}}

	res := ""

	for _, f := range factors {
		if n%f.factor == 0 {
			res += f.name
		}
	}

	if res == "" {
		res = itoa(n)
	}

	return res
}
