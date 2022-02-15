package adapter

import (
	"crypto/md5"
	"fmt"

	"github.com/kinix/http-body-hash-generator/app"
)

type MD5Generator struct {
}

// Validate interface compliance
var _ app.HashGenerator = (*MD5Generator)(nil)

// Generate MD5 as a string
func (g *MD5Generator) Generate(input string) string {
	val := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", val)
}
