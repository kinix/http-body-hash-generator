package adapter_test

import (
	"testing"

	"github.com/kinix/http-body-hash-generator/adapter"
)

func TestMD5Generator(t *testing.T) {
	type testCase struct {
		input  string
		output string
	}

	testCases := []testCase{
		{"hash me", "17b31dce96b9d6c6d0a6ba95f47796fb"},
		{"something", "437b930db84b8079c2dd804a71936b5f"},
		{"something else", "6c7ba9c5a141421e1c03cb9807c97c74"},
	}

	generator := &adapter.MD5Generator{}

	for _, test := range testCases {
		result := generator.Generate(test.input)
		if result != test.output {
			t.Errorf("%s is expected, but %s is returned for %s", test.output, result, test.input)
		}
	}
}
