// Reverse-string task of Exercism.io Go track.
package reverse

// Reverse string. Simple and slow.
func Reverse1(s string) string {
	res := ""
	for _, r := range s {
		res = string(r) + res
	}
	return res
}

// Reverse string with some optimization in mind.
func Reverse(s string) string {
	runes := make([]rune, len(s))
	count := 0

	// Fill array of runes.
	for _, r := range s {
		runes[count] = r
		count++
	}

	// 'count' now holds number of runes in the string.
	// 'count' <= 'len(s)' as single rune may take more than one byte.

	// Reverse array of runes.
	h, t := 0, count-1
	for h < t {
		runes[h], runes[t] = runes[t], runes[h]
		h++
		t--
	}
	return string(runes[:count])
}
