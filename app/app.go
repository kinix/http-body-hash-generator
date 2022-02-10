package app

// Adapter list
type App struct {
	config        ConfigParser
	httpClient    HttpClient
	hashGenerator HashGenerator
	outputWriter  OutputWriter
}

// Parser for URL list and parallelJobCount from CLI, file or another resource
type ConfigParser interface {
	GetList() []string
	GetParallelJobCount() int
}

// HTTP Client (which includes a function to send GET requests and read full body)
type HttpClient interface {
	GetBody(string) (string, error)
}

// Generate hash as string (might be MD5, SHA256 etc.)
type HashGenerator interface {
	Generate(string) string
}

// Printer for write output to CLI, file etc.
type OutputWriter interface {
	Write(string, string)
}

// Create app and set the adapters
func NewApp(config ConfigParser, httpClient HttpClient, hashGenerator HashGenerator, outputWriter OutputWriter) *App {
	app := &App{}

	app.config = config
	app.httpClient = httpClient
	app.hashGenerator = hashGenerator
	app.outputWriter = outputWriter

	return app
}

func (a *App) Run() {
	urlList := a.config.GetList()
	workerCount := a.config.GetParallelJobCount()

	// URL channel will be used to send URLs to workers
	urlChannel := make(chan string, len(urlList))

	// Finished channel will be used to be sure that the worker is finihed its job
	finished := make(chan bool, len(urlList))

	// Create workers
	for i := 0; i < workerCount; i++ {
		go a.worker(urlChannel, finished)
	}

	// Send URLs
	for _, url := range urlList {
		urlChannel <- url
	}

	// After all URLs is sended, close the channel
	close(urlChannel)

	// Receive finished signals
	// When the channel received one signal for each URL, it means all jobs are finished
	for i := 0; i < len(urlList); i++ {
		<-finished
	}

	// Close the channel
	close(finished)
}

func (a *App) worker(urlChannel <-chan string, finished chan<- bool) {
	// Whenever channel sends a URL, one of worker get it
	for url := range urlChannel {
		// Send HTTP request and read the body
		body, err := a.httpClient.GetBody(url)

		// If there is an error in the request, print it
		if err != nil {
			a.outputWriter.Write(url, err.Error())
			continue
		}

		// Hash the body and print it
		hash := a.hashGenerator.Generate(body)
		a.outputWriter.Write(url, hash)

		// Send a signal about finished job
		finished <- true
	}
}
