package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// testingT defines the api that is used from testing.T
// this exists for testing Assert using a mock.
type testingT interface {
	Errorf(string, ...interface{})
	Log(args ...interface{})
}

// testingA decorates a *testing.T to create an Operand using That(..) and do error logging
type testingA struct {
	t testingT
}

// That creates an Operand on the value we have got and describes the variable that is being testing.
func (a testingA) That(label string, got interface{}) Operand {
	return Operand{a, label, got, equals{}}
}

// That creates an Operand on the value we have got and describes the variable that is being testing.
func That(t testingT, label string, got interface{}) Operand {
	return Operand{testingA{t}, label, got, equals{}}
}

// Assert creates an Operand on a value that needs to be checked.
func Assert(t testingT, label string, value interface{}) Operand {
	return testingA{t}.That(label, value)
}

// Asser is more a syntax trick to shorten the amount of code needed to create an Operand.
// So instead of:
//		assert.That("age",age).Equals(42)
// or
//  	Assert(t,"age",age).Equals(42)
// you can write:
//		Asser(t,"age",age).Equals(42)
var Asser = Assert
