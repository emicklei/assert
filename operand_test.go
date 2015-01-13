package assert

import (
	"strings"
	"testing"
)

func TestIntEqualsInt(t *testing.T) {
	testingA{t}.That("i", 10).Equals(10)
}

func TestStringEqualsString(t *testing.T) {
	testingA{t}.That("s", "abcd").Equals("abcd")
}

func TestInt32EqualsInt(t *testing.T) {
	testingA{t}.That("int", int32(32)).Equals(32)
}

func TestStringEqualsInt(t *testing.T) {
	r := new(testReporter)
	testingA{r}.That("string", "string").Equals(32)
	if len(r.args) == 0 {
		t.Fail()
	}
	// tricky, we use Assert to test an assert feature
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
	Asser(t, "bool", true).IsTrue()
}

func TestNot(t *testing.T) {
	Asser(t, "bool", false).Not().IsTrue()
}

func TestNot_Fail(t *testing.T) {
	r := new(testReporter)
	Asser(r, "bool", true).Not().IsTrue()
	if len(r.args) == 0 {
		t.Fail()
	}
}

func TestIsNil(t *testing.T) {
	var n error
	Asser(t, "nil", n).IsNil()
}

func TestIsNil_Fail(t *testing.T) {
	var n error
	r := new(testReporter)
	Asser(r, "nil", n).Not().IsNil()
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

func TestCompareUsing(t *testing.T) {
	Asser(t, "insensitive", "ABC").CompareUsing(caseInsensitiveStringEquals{}).Equals("abc")
}

type caseInsensitiveStringEquals struct{}

func (c caseInsensitiveStringEquals) Compare(left, right interface{}) bool {
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
