// package bob implements Exercism task 'bob'.
package bob

import (
	"strings"
	"unicode"
)

// isQuestion returns true if given remark ends with a question and returns
// false otherwise.
func isQuestion(remark string) bool {
	l := len(remark)
	return l > 0 && remark[l-1] == '?'
}

// isYelling returns true if remark contains some letters and all of them are
// in upper case.
// It returns false otherwise.
func isYelling(remark string) bool {
	letters := 0
	for _, r := range remark {
		if unicode.IsLetter(r) {
			letters++
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return letters > 0
}

// isBlank returns true if s does not contain letters or digits.
// It returns false otherwise.
func isBlank(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// testF is a function that tests string argument to match some condition.
type testF func(string) bool

// composeAnd retunrs funcion that is AND composition of f and optional list
// of functions.
func composeAnd(f testF, rest ...testF) testF {
	return func(s string) bool {
		res := f(s)
		for i := 0; i < len(rest) && res; i++ {
			res = rest[i](s)
		}
		return res
	}
}

// Hey returns a string which is Bob's response to given remark.
// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying
// anything.
// He answers 'Whatever.' to anything else.
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r\v\b")

	bob := []struct {
		match  testF
		answer string
	}{
		// IMPORTANT: order does matter
		{composeAnd(isQuestion, isYelling), "Calm down, I know what I'm doing!"},
		{isQuestion, "Sure."},
		{isYelling, "Whoa, chill out!"},
		{isBlank, "Fine. Be that way!"}}

	for _, b := range bob {
		if b.match(remark) {
			return b.answer
		}
	}

	return "Whatever."
}
