package assert

import "testing"

func TestIntEqualsInt(t *testing.T) {
	Assert{t}.That("i", 10).Equals(10)
}

func TestStringEqualsString(t *testing.T) {
	Assert{t}.That("s", "abcd").Equals("abcd")
}

func TestInt32EqualsInt(t *testing.T) {
	Assert{t}.That("int", int32(32)).Equals(32)
}

func TestStringEqualsInt(t *testing.T) {
	r := new(testReporter)
	Assert{r}.That("string", "string").Equals(32)
	if len(r.args) == 0 {
		t.Fail()
	}
	// tricky, we use Assert to test an assert feature
	Assert{t}.That("arg.1", r.args[2]).Equals("string")
}

func TestStringIsKindOfString(t *testing.T) {
	Assert{t}.That("string", "string").IsKindOf("other")
}

func TestStringIsKindOfInt(t *testing.T) {
	r := new(testReporter)
	Assert{r}.That("string", "string").IsKindOf(42)
	if len(r.args) == 0 {
		t.Fail()
	}
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

//// clear && go test -v -test.run TestStringEqualsString_Fail ...script
//func TestStringEqualsString_Fail(t *testing.T) {
//	Assert{t}.That("s", "abcd").Equals("abc")
//}

//// clear && go test -v -test.run TestIntEqualsInt_Fail ...script
//func TestIntEqualsInt_Fail(t *testing.T) {
//	Assert{t}.That("i", 24).Equals(10)
//}

//// clear && go test -v -test.run TestIntEqualsInt_WrongType ...script
//func TestIntEqualsInt_WrongType(t *testing.T) {
//	Assert{t}.That("i", 1.0).Equals(1)
//}
