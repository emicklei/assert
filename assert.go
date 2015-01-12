package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// testingT defines the api that is used from testing.T
// this exists for testing Assert using a mock.
type testingT interface {
	Errorf(string, ...interface{})
}

// Assert decorates a *testing.T to create an Operand using That(..)
type Assert struct {
	t testingT
}

// That creates an Operand on the value we have got and describs the variable that is being testing.
func (a Assert) That(label string, got interface{}) Operand {
	return Operand{a, label, got, EqualsComparator{}}
}

// That creates an Operand on the value we have got and describs the variable that is being testing.
func That(t testingT, label string, got interface{}) Operand {
	return Operand{Assert{t}, label, got, EqualsComparator{}}
}

// Asser is a syntax trick to shorten the amount of code needed to create an Operand.
// So instead of:
//	assert.That("age",age).Equals(42)
// you can write:
//	Asser(t,"age",age).Equals(42)
func Asser(t testingT, label string, value interface{}) Operand {
	return Assert{t}.That(label, value)
}
