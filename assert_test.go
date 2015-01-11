package assert

type testReporter struct {
	template string
	args     []interface{}
}

func (r *testReporter) Errorf(template string, args ...interface{}) {
	r.template = template
	r.args = args
}
