package adapter

import (
	"crypto/md5"
	"fmt"
)

type MD5Generator struct {
}

// Generate MD5 as a string
func (g *MD5Generator) Generate(input string) string {
	val := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", val)
}
