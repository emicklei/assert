/*
Package assert, for writing checks in unit tests

This package provides functions to reduce the amount of code needed to write simple assertions.
It implements the best practice pattern where the output of a failure explains what the check "got" and what it wanted.
The assert functions are defined such that writing requires less code but still are easy to understand.
It works by decorating the standard testing.T type in your test and report (Fatal) the offending asserting call if a check fails.

Example

	import (
		"testing.T"
		"github.com/emicklei/assert"
	)

	func TestShoeFits(t *testing.T) {
		shoeSize := 42
		assert.That(t,"shoeSize",shoeSize).Equals(46)
	}

which will report

	got [42] (int) for "shoeSize" but want [46] (int)


Examples: (using the dot import)

	Asser(t,"err",err).IsNil()
	Asser(t,"isOffline",isOffline).IsTrue()
	Asser(t,"country",country).Equals("NL")
	Asser(t,"job",job).IsKindOf(new(Job))
	Asser(t,"names", []string{}).Len(0)

	// you can negate a check
	Asser(t,"isOnline",isOnline).Not().IsTrue()
*/
package assert
