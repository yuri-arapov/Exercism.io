// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// type Year represents a year in Gregorian calendar
type Year int

// divisibleBy returns true if y is evenly divisible by divider
func (y Year) divisibleBy(divider int) bool {
	return (int(y) % divider) == 0
}

// IsLeap returns true if y is a leap year
func (y Year) IsLeap() bool {
	// a leap year in the Gregorian calendar occurs:
	// on every year that is evenly divisible by 4
	//  except every year that is evenly divisible by 100
	//    unless the year is also evenly divisible by 400
	return (y.divisibleBy(4) && !y.divisibleBy(100)) || y.divisibleBy(400)
}

// IsLeapYear returns true if y is a leap year
func IsLeapYear(y int) bool {
	return Year(y).IsLeap()
}
