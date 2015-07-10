package assert

import (
	"fmt"

	"github.com/wsxiaoys/terminal/color"
)

// TerminalColorsEnabled can be changed to disable the use of terminal coloring.
// One usecase is to add a command line flag to your test that controls its value.
//
//	func init() {
//		flag.BoolVar(&assert.TerminalColorsEnabled, "color", true, "want colors?")
//	}
//
//	go test -color=false
var TerminalColorsEnabled = true

// FatalColorSyntaxCode requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go .
// Set to an empty string to disable coloring.
var FatalColorSyntaxCode = "@{wB}"

// SuccessColorSyntaxCode requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go .
// Set to an empty string to disable coloring.
var SuccessColorSyntaxCode = "@{y}"

// Scolorf returns a string colorized for terminal output using the syntaxCode (unless that's empty).
// Requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go .
func Scolorf(syntaxCode string, format string, args ...interface{}) string {
	plainFormatted := fmt.Sprintf(format, args...)
	if len(syntaxCode) > 0 && TerminalColorsEnabled {
		// cannot pass the code as a string param
		return color.Sprintf(syntaxCode+"%s", plainFormatted)
	}
	return plainFormatted
}
