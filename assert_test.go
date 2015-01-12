package assert

// clear && go test -coverprofile=coverage.out && go tool cover -func=coverage.out

type testReporter struct {
	template string
	args     []interface{}
}

func (r *testReporter) Errorf(template string, args ...interface{}) {
	r.template = template
	r.args = args
}
