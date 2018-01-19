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
	"fmt"
	"os"
	"strconv"
	"path/filepath"
	"strings"

	"gopkg.in/urfave/cli.v1"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	var url string
	var code string
	var method string
	var schemaFilename string
	var headersFilename string
	var bodyFilename string
	var silent bool

	app := cli.NewApp()
	app.Name = "verify-url"
	app.Usage = "Basic testing for URL responses"
	app.UsageText = "verify-url [options] [URL]";
	app.Version = "1.0.0";
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "code, c",
			Value: "200",
			Usage: "The expected HTTP status code. Defaults to 200.",
			Destination: &code,
		},
		cli.StringFlag{
			Name: "method, m",
			Value: "GET",
			Usage: "The HTTP method to use when calling the URL. Defaults to 'GET'.",
			Destination: &method,
		},
		cli.StringFlag{
			Name: "schema, sch",
			Usage: "`FILE` used to load JSON schema for response verification",
			Destination: &schemaFilename,
		},
		cli.StringFlag{
			Name: "headers, hd",
			Usage: "`FILE` containing headers to send to URL",
			Destination: &headersFilename,
		},
		cli.StringFlag{
			Name: "body, b",
			Usage: "`FILE` containing body content to send to URL",
			Destination: &bodyFilename,
		},
		cli.BoolFlag{
			Name: "silent, s",
			Usage: "If specified, nothing will be printed to stdout",
			Destination: &silent,
		},
	}

	app.Action = func(c *cli.Context) error {
		url = c.Args().Get(0)
		if url == "" {
			return HandleError("URL is required.\n", silent)
		}

		parsedCode, err := strconv.ParseInt(code, 0, 0)
		if err != nil {
			return HandleError("Could not parse the provided status code.\n", silent)
		}

		body, status, err := makeRequest(url)
		if err != nil {
			return HandleError(fmt.Sprintf("An error was encountered: \"%s\".\n", err.Error()), silent)
		}

		if status != int(parsedCode) {
			return HandleError("Actual status code does not match expected status code.\n", silent)
		}

		if schemaFilename != "" {
			schemaFilename, err := filepath.Abs(schemaFilename)
			if err != nil {
				return HandleError("Could not use provided schema file name.\n", silent)
			}

			schemaFilename = URIFromPath(schemaFilename)

			loadedSchema := gojsonschema.NewReferenceLoader(schemaFilename)
			loadedBody := gojsonschema.NewStringLoader(body)
			result, err := gojsonschema.Validate(loadedSchema, loadedBody)
			if err != nil {
				return HandleError("Validation error: " + err.Error() + "\n", silent)
			}

			if !result.Valid() {
				return HandleError("Response did not match the provided schema.\n", silent)
			}
		}

		if !silent {
			fmt.Fprintf(os.Stdout, "%s\n", body)
		}

		return nil;
	}

	app.Run(os.Args)
}

func HandleError(message string, silent bool) error {
	if silent {
		return cli.NewExitError("", 1)
	}

	return cli.NewExitError(message, 1)
}

func URIFromPath(path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "file://") {
		return path
	}

	return "file://" + path;
}