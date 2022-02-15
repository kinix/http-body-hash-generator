package adapter

import (
	"fmt"

	"github.com/kinix/http-body-hash-generator/app"
)

type CLIPrinter struct {
}

// Validate interface compliance
var _ app.OutputWriter = (*CLIPrinter)(nil)

// Print url and the result to command line interface
func (c *CLIPrinter) Write(url string, result string) {
	fmt.Printf("%s %s\n", url, result)
}
