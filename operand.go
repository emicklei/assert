package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import "reflect"

// Operand represent a value
type Operand struct {
	// this reference is used to report a test failure
	a testingA
	// description of the value, typically a variable or field name
	label string
	// actual value of any type
	value interface{}
	// used to operate on two values
	operator RelationalOperator
}

// OperateUsing returns a copy Operand that will use the RelationalOperator.
func (o Operand) OperateUsing(r RelationalOperator) Operand {
	return Operand{o.a, o.label, o.value, r}
}

// Equals checks whether the value we have got is equal to the value we want.
func (o Operand) Equals(want interface{}) {
	not := not{o.operator}
	if not.Apply(o.value, want) {
		// if the right value (what we want) has int type
		// then try convert the right value using reflection and re-compare
		// See https://golang.org/ref/spec#Numeric_types
		_, ok := want.(int)
		if ok {
			leftType := reflect.TypeOf(o.value)
			if leftType != nil { // if left is nil
				rightValue := reflect.ValueOf(want)
				convertedRightValue := rightValue.Convert(leftType)
				if reflect.ValueOf(o.value) == convertedRightValue {
					return
				}
			}
		}
		logCall(o.a.t, "Equals")
		o.a.t.Errorf("got [%v] (%T) for \"%s\" but want [%v] (%T)",
			o.value, o.value,
			o.label,
			want, want)
	}
}

// IsKindOf checks whether the values are of the same type
func (o Operand) IsKindOf(v interface{}) {
	leftType := reflect.TypeOf(o.value)
	rightType := reflect.TypeOf(v)
	if leftType != rightType {
		logCall(o.a.t, "IsKindOf")
		o.a.t.Errorf("got [%v] for \"%s\" but want [%v]", leftType, o.label, rightType)
	}
}

// IsNil checks whether the value is nil
func (o Operand) IsNil() {
	if !o.operator.Apply(o.value, nil) {
		logCall(o.a.t, "IsNil")
		o.a.t.Errorf("got [%v] for \"%s\" but want [nil]", o.value, o.label)
	}
}

// IsNotNil checks whether the value is nil
func (o Operand) IsNotNil() {
	if o.operator.Apply(o.value, nil) {
		logCall(o.a.t, "IsNotNil")
		o.a.t.Errorf("got unexpected [%v] for \"%s\"", o.value, o.label)
	}
}

// IsTrue checks whether the value is true
func (o Operand) IsTrue() {
	if o.operator.Apply(o.value, false) { // i.e fail if !true
		logCall(o.a.t, "IsTrue")
		o.a.t.Errorf("got [%v] for \"%s\" but want [true]", o.value, o.label)
	}
}

// IsFalse checks whether the value is false
func (o Operand) IsFalse() {
	if o.operator.Apply(o.value, true) { // i.e fail if !false
		logCall(o.a.t, "IsFalse")
		o.a.t.Errorf("got [%v] for \"%s\" but want [false]", o.value, o.label)
	}
}

// Not creates a new Operand with a negated version of its comparator.
func (o Operand) Not() Operand {
	return Operand{o.a, o.label, o.value, not{o.operator}}
}
