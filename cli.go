/*
Copyright 2017 Toolhouse, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"gopkg.in/urfave/cli.v1"
)

// Args contains a set of command line arguments
type Args struct {
	code            uint
	method          string
	schemaFilename  string
	headersFilename string
	bodyFilename    string
	silent          bool
}

func setupCli(app *cli.App, version string) *Args {
	info := Args{}

	app.Name = "validate-http-response"
	app.Usage = "Basic testing for HTTP responses"
	app.UsageText = "validate-http-response [options] [URL]"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.UintFlag{
			Name:        "code",
			Value:       200,
			Usage:       "The expected HTTP status code. Defaults to 200.",
			Destination: &info.code,
		},
		cli.StringFlag{
			Name:        "method",
			Value:       "GET",
			Usage:       "The HTTP method to use when calling the URL. Defaults to 'GET'.",
			Destination: &info.method,
		},
		cli.StringFlag{
			Name:        "schema",
			Usage:       "`FILE` used to load JSON schema for response verification",
			Destination: &info.schemaFilename,
		},
		cli.StringFlag{
			Name:        "headers",
			Usage:       "`FILE` containing headers to send to URL",
			Destination: &info.headersFilename,
		},
		cli.StringFlag{
			Name:        "body",
			Usage:       "`FILE` containing body content to send to URL",
			Destination: &info.bodyFilename,
		},
		cli.BoolFlag{
			Name:        "silent",
			Usage:       "If specified, nothing will be printed to stdout",
			Destination: &info.silent,
		},
	}

	return &info
}
