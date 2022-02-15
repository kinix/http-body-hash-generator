package adapter

import (
	"testing"
)

func TestArgParser(t *testing.T) {
	type testCase struct {
		args             []string
		list             []string
		parallelJobCount int
	}

	testCases := []testCase{
		{
			[]string{"adjust.com", "google.com", "reddit.com/r/funny"},
			[]string{"adjust.com", "google.com", "reddit.com/r/funny"},
			defaultParallelJobCount,
		},
		{
			[]string{"-parallel", "42", "adjust.com", "google.com", "reddit.com/r/funny"},
			[]string{"adjust.com", "google.com", "reddit.com/r/funny"},
			42,
		},
		{
			[]string{"adjust.com", "google.com", "reddit.com/r/funny", "-parallel", "42"},
			[]string{"adjust.com", "google.com", "reddit.com/r/funny"},
			42,
		},
		{
			[]string{"-parallel", "adjust.com", "google.com", "reddit.com/r/funny"},
			[]string{"google.com", "reddit.com/r/funny"},
			defaultParallelJobCount,
		},
	}

	for _, test := range testCases {
		parser := NewArgParser()

		// New parser parses CLI args default, so we need to override args
		parser.ParseArgList(test.args)

		list := parser.GetList()
		for i, val := range list {
			// URL list should contain the same items
			if val != test.list[i] {
				t.Errorf("List member %d is different from the expected value. Expected: %s, Returned: %s",
					i, test.list[i], val)
			}
		}

		// parallelJobCount should be the same
		if parser.GetParallelJobCount() != test.parallelJobCount {
			t.Errorf("Parallel job count is different from the expected value. Expected: %d, Returned: %d",
				test.parallelJobCount, parser.GetParallelJobCount())
		}
	}
}
