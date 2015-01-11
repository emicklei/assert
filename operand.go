package assert

import "reflect"

// Operand represent a value
type Operand struct {
	// this reference is used to report a test failure
	a Assert
	// description of the value, typically a variable or field name
	label string
	// actual value of any type
	value interface{}
	// used to compare two values
	comparator Comparator
}

// Equals checks whether the value we have got is equal to the value we want.
func (o Operand) Equals(want interface{}) {
	not := Not{o.comparator}
	if not.Compare(o.value, want) {
		// if the right value (what we want) has int type
		// then try convert the right value using reflection and re-compare
		// See https://golang.org/ref/spec#Numeric_types
		_, ok := want.(int)
		if ok {
			leftType := reflect.TypeOf(o.value)
			rightValue := reflect.ValueOf(want)
			convertedRightValue := rightValue.Convert(leftType)
			if reflect.ValueOf(o.value) == convertedRightValue {
				return
			}
		}
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
		o.a.t.Errorf("got [%v] for \"%s\" but want [%v]", leftType, o.label, rightType)
	}
}

// IsNil checks whether the value is nil
// Use Not().IsNil() to check that the value is not nil
func (o Operand) IsNil() {
	if o.value != nil {
		o.a.t.Errorf("got [%v] for \"%s\"", o.value, o.label)
	}
}

// IsTrue checks whether the value is true
// Use Not().IsTrue() to check against false
func (o Operand) IsTrue() {
	if o.comparator.Compare(o.value, false) {
		o.a.t.Errorf("got [%v] for \"%s\" but want [false]", o.value, o.label)
	}
}

// Not creates a new Operand with a negated version of its comparator.
func (o Operand) Not() Operand {
	return Operand{o.a, o.label, o.value, Not{o.comparator}}
}
