// Collatz Conjecture problem of Exercism.io Go Track.
package collatzconjecture

import "fmt"

// Return number of steps needed to converge Collatz Conjecture from n to 1.
// Error is not nil if n is not positive.
func CollatzConjecture(n int) (int, error) {
	//fmt.Printf("%d", n)
	if n <= 0 {
		//fmt.Printf("<-invalid argument\n")
		return 0, fmt.Errorf("invalid argument %d", n)
	}
	s := 0
	for ; n != 1; s++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	//fmt.Printf("->%d\n", s)
	return s, nil
}
