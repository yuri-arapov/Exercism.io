//package main
package listops

import (
	"fmt"
)

func main() {
	var s *IntList = nil
	fmt.Printf("Length(%v) = %d\n", s, s.Length())

	s = MakeList(1, 2, 3)
	fmt.Printf("Length(%v) = %d\n", s, s.Length())

	var n *IntList
	fmt.Printf("Append: %v\n", n.Append(nil))

	fmt.Printf("Append: %v\n", s.Append(MakeList(4, 5, 6)))

	tt := MakeList(4, 5, 6, 7, 8)
	fmt.Printf("Tail: %v %v\n", tt, tt.Tail())

	fmt.Printf("Reverse: %v %v\n", tt, tt.Reverse())

	fmt.Printf("Filter: %v %v\n", tt, tt.Filter(func(n int) bool { return n&1 != 0 }))
	fmt.Printf("Filter: %v %v\n", tt, tt.Filter(func(n int) bool { return n&1 == 0 }))

	fmt.Printf("Map: %v %v\n", tt, tt.Map(func(n int) int { return n * 4 }))
}
