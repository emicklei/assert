package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// clear && go test -coverprofile=coverage.out && go tool cover -func=coverage.out

// clear &&  go tool cover -func=coverage.out && go tool cover -html=coverage.out

type testReporter struct {
	template string
	args     []interface{}
}

func (r *testReporter) Fatalf(template string, args ...interface{}) {
	r.template = template
	r.args = args
}

func (r *testReporter) Log(args ...interface{}) {}

func (r *testReporter) Helper() {}

func (r *testReporter) Logf(format string, args ...interface{}) {}

func newTestReporter() *testReporter {
	t := new(testReporter)
	// capture fatals
	Fatalf = func(t testingT, format string, args ...interface{}) {
		t.Fatalf(format, args...)
	}
	return t
}
