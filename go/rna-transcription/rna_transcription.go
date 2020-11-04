// RNA transcription exercise of Exercism.io Go track.
package strand

// Convern DNA strand to RNA.
func ToRNA(dna string) string {
	// `G` -> `C`
	// `C` -> `G`
	// `T` -> `A`
	// `A` -> `U`
	dna2rna := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	rna := make([]rune, len(dna))
	for i, n := range dna {
		rna[i] = dna2rna[n]
	}
	return string(rna)
}
