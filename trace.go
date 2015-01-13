package assert

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func logCall(t testingT, funcName string) {
	var buffer bytes.Buffer
	insidePkg := true
	lastLine := 0
	lastFile := ""
	for i := 0; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !insidePkg {
			buffer.WriteString(fmt.Sprintf("%s() call at %s:%d", funcName, lastFile, lastLine))
			break
		}
		if strings.Index(file, "github.com/emicklei/assert") == -1 {
			insidePkg = false
		}
		lastLine = line
		lastFile = file

	}
	t.Log(buffer.String())
}
