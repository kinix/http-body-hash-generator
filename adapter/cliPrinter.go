package adapter

import "fmt"

type CLIPrinter struct {
}

// Print url and the result to command line interface
func (c *CLIPrinter) Write(url string, result string) {
	fmt.Printf("%s %s\n", url, result)
}
