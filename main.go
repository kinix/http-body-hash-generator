package main

import (
	"github.com/kinix/http-body-hash-generator/adapter"
	"github.com/kinix/http-body-hash-generator/app"
)

func main() {
	// Create adapters
	var parser app.ConfigParser
	var httpClient app.HttpClient
	var md5generator app.HashGenerator
	var cliPrinter app.OutputWriter

	parser = adapter.NewArgParser()
	httpClient = &adapter.HttpClient{}
	md5generator = &adapter.MD5Generator{}
	cliPrinter = &adapter.CLIPrinter{}

	// Run the app
	baseApp := app.NewApp(parser, httpClient, md5generator, cliPrinter)
	baseApp.Run()
}
