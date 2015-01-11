package assert

// Comparator specifies the function to compare two values
type Comparator interface {
	// Return the result of comparing left and right
	Compare(left, right interface{}) bool
}

// EqualsComparator can compare two values using ==
type EqualsComparator struct{}

// Compare returns the result of comparing left and right using ==
func (c EqualsComparator) Compare(left, right interface{}) bool {
	return left == right
}

// Not is to negate the result of a Comparator
type Not struct {
	c Comparator
}

// Compare returns the opposite boolean result of comparing left and right
func (n Not) Compare(left, right interface{}) bool {
	return !n.c.Compare(left, right)
}
