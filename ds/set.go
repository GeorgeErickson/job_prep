package ds

// Set is a hash based set.
type Set map[interface{}]struct{}

// Len returns number of items in set
func (s *Set) Len() int {
	return len(*s)
}

// Add idempotently adds item to set.
func (s *Set) Add(v interface{}) {
	(*s)[v] = struct{}{}
}

// Contains returns whether the set contains the given element.
func (s *Set) Contains(v interface{}) bool {
	_, exists := (*s)[v]
	return exists
}

// Remove `v` from the set.
func (s *Set) Remove(v interface{}) {
	delete(*s, v)
}

// Each calls `fn` for each item in the set.
func (s *Set) Each(fn func(v interface{})) {
	for v := range *s {
		fn(v)
	}
}
