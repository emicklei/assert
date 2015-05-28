package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"errors"
	"strings"
	"testing"
)

func TestAbsentMapValueEqualsInt(t *testing.T) {
	r := new(testReporter)
	var data map[string]interface{}
	testingA{r}.That("", data["missing"]).Equals(1)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestIntEqualsInt(t *testing.T) {
	testingA{t}.That("i", 10).Equals(10)
}

func TestStringEqualsString(t *testing.T) {
	testingA{t}.That("s", "abcd").Equals("abcd")
}

func TestStringEqualsInt(t *testing.T) {
	r := newTestReporter()
	testingA{r}.That("string", "string").Equals(32)
	if len(r.args) == 0 {
		t.Fail()
	}
	// tricky, we use Assertt to test an assert feature
	testingA{t}.That("arg.1", r.args[2]).Equals("string")
}

func TestStringIsKindOfString(t *testing.T) {
	testingA{t}.That("string", "string").IsKindOf("other")
}

func TestStringIsKindOfInt(t *testing.T) {
	r := new(testReporter)
	testingA{r}.That("string", "string").IsKindOf(42)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestIsTrue(t *testing.T) {
	Assert(t, "bool", true).IsTrue()
}

func TestIsFalse(t *testing.T) {
	Assert(t, "bool", false).IsFalse()
}

func TestIsFalse_Fail(t *testing.T) {
	r := new(testReporter)
	Assert(r, "bool", true).IsFalse()
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestNot(t *testing.T) {
	Assert(t, "bool", false).Not().IsTrue()
}

func TestNot_Fail(t *testing.T) {
	r := new(testReporter)
	Assert(r, "bool", true).Not().IsTrue()
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestIsNil(t *testing.T) {
	var n error
	Assert(t, "nil", n).IsNil()
}

func TestIsNil_Fail(t *testing.T) {
	var n error
	r := new(testReporter)
	Assert(r, "nil", n).Not().IsNil()
	if len(r.template) == 0 {
		t.Fail()
	}
}

func TestIsNotNil(t *testing.T) {
	nn := errors.New("not")
	Assert(t, "notnil", nn).IsNotNil()
}

func TestIsNotNil_Fail(t *testing.T) {
	var nn error
	r := new(testReporter)
	Assert(r, "notnil", nn).IsNotNil()
	if len(r.template) == 0 {
		t.Fail()
	}
}

func TestNotNot(t *testing.T) {
	That(t, "not-not", true).Not().Not().IsTrue()
}

func TestStringEqualsString_Fail(t *testing.T) {
	r := new(testReporter)
	testingA{r}.That("s", "abcd").Equals("abc")
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestIntEqualsInt_Fail(t *testing.T) {
	r := new(testReporter)
	testingA{r}.That("i", 24).Equals(10)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestLogCall(t *testing.T) {
	doLogCall(t)
}

func doLogCall(t *testing.T) {
	logCall(t, "TestLogCall")
}

func TestCompareUsing(t *testing.T) {
	Assert(t, "insensitive", "ABC").With(caseInsensitiveStringEquals{}).Equals("abc")
}

type caseInsensitiveStringEquals struct{}

func (c caseInsensitiveStringEquals) Apply(left, right interface{}) bool {
	s_left, ok := left.(string)
	if !ok {
		return false
	}
	s_right, ok := right.(string)
	if !ok {
		return false
	}
	return strings.EqualFold(s_left, s_right)
}

type understandsLen struct{}

func (u understandsLen) Len() int { return 0 }

func TestLen(t *testing.T) {
	list := []string{}
	Assert(t, "list", list).Len(0)
	Assert(t, "list", " ").Len(1)
	Assert(t, "list", map[string]int{}).Len(0)
	ch := make(chan int)
	Assert(t, "chan", ch).Len(0)
	Assert(t, "custom", understandsLen{}).Len(0)
}

func TestLen_FailSlice(t *testing.T) {
	r := new(testReporter)
	list := []string{}
	testingA{r}.That("list", list).Len(1)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestLen_FailNoLen(t *testing.T) {
	r := new(testReporter)
	testingA{r}.That("reporter", r).Len(0)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestLen_FailWrongLen(t *testing.T) {
	r := new(testReporter)
	custom := understandsLen{}
	testingA{r}.That("custom", custom).Len(42)
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestSliceEqualsInt_Fail(t *testing.T) {
	r := new(testReporter)
	list := []string{}
	testingA{r}.That("list", list).Equals(1)
	if len(r.args) == 0 {
		t.Fail()
	}
	//t.Logf(r.template, r.args...)
}
