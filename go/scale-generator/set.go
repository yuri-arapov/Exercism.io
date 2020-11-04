package scale

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
