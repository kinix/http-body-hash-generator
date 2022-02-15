package adapter

import (
	"log"
	"os"
	"strconv"

	"github.com/kinix/http-body-hash-generator/app"
)

const (
	ParallelJobArg          = "-parallel"
	DefaultParallelJobCount = 10
)

// The object to parse arguments for configs
type ArgParser struct {
	parallelJobCount int
	list             []string
}

// Validate interface compliance
var _ app.ConfigParser = (*ArgParser)(nil)

func NewArgParser() *ArgParser {
	parser := &ArgParser{}

	// Read and parse CLI args
	args := getArgList()
	parser.ParseArgList(args)

	return parser
}

func getArgList() []string {
	// Ignore the first argument because it is the app name
	return os.Args[1:]
}

// Parser parses CLI args, but ParseArgList can be used if you need to parse a custom string
func (a *ArgParser) ParseArgList(args []string) {
	// Clear the list
	a.list = []string{}

	parallel := false

	for _, arg := range args {
		// If the previous arg was the reserved word for parallel jobs
		if parallel {
			// The arg should be a positive number
			if parallelJobCount, err := strconv.Atoi(arg); err != nil || parallelJobCount < 1 {
				log.Println("Parallel argument is not a positive number")
			} else {
				a.parallelJobCount = parallelJobCount
			}

			parallel = false
			continue
		}

		// Check if the arg is reserved word for parallel jobs
		if arg == ParallelJobArg {
			parallel = true
			continue
		}

		// If it is not about parallel jobs, add the arg to the list
		a.list = append(a.list, arg)
	}

	// If there is no arg for parallel jobs, set it as the default value
	if a.parallelJobCount == 0 {
		a.parallelJobCount = DefaultParallelJobCount
	}
}

// Return the list
func (a *ArgParser) GetList() []string {
	return a.list
}

// Return the parallel job count
func (a *ArgParser) GetParallelJobCount() int {
	return a.parallelJobCount
}
