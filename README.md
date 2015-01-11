Assert is yet another package for writing unit tests in Go.

Reason for creating it:
- code for unit testing must be "close" to the standard ; no Suites, DSL or other functional weirdness.
- if a check fails then try hard to explain why and what went wrong
- allow checks of all integer types against int, e.g. int32(42) == 42
- allow for custom comparators
- compact but still readable

### Example

	import (
		"testing.T"
		. "github.com/emicklei/assert"
	)

	func TestShoeFits(t *testing.T) {
		shoeSize := 42
		Asser(t,"shoeSize",shoeSize).Equals(46)
	}
	
which will report

	got [46] (int) for "shoeSize" but want [42] (int)
	

### Snippets

	Asser(t,"err",err).IsNil()
	Asser(t,"isOffline",isOffline).IsTrue()
	Asser(t,"country",country).Equals("NL")
	Asser(t,"job",job).IsKindOf(new(Job))
	
	Asser(t,"isOnline",isOnline).Not().IsTrue()
	Asser(t,"policy",policy).Not().IsNil()
		
or without using the dot import:

	assert.That(t,"variable",got).Equals(expected)
