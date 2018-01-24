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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const (
	validatorErrorMissingUrl             = "URL is required."
	validatorErrorHeaderFileReadError    = "Could not read the provided headers file."
	validatorErrorHeaderDeserializeError = "Could not deserialize headers."
	validatorErrorBodyFileReadError      = "Could not read the provided body file."
	validatorErrorHttpRequestError       = "HTTP error: \"%s\"."
	validatorErrorStatusCodeMismatch     = "Actual status code (%d) does not match expected status code (%d)."
	validatorErrorSchemaFileError        = "Could not use provided schema file name."
	validatorErrorGeneric                = "Validation error: \"%s\"."
	validatorErrorSchemaMismatch         = "Response did not match the provided schema."
)

// Validator validates a URL against a set of command line arguments
type Validator struct {
	args *Args
}

// Run validator on the specified URL
func (v Validator) Run(url string) error {
	if url == "" {
		return v.handleError(validatorErrorMissingUrl)
	}

	response, err := v.makeRequest(url)
	if err != nil {
		return err
	}

	err = v.validateResponse(response)
	if err != nil {
		return err
	}

	if !v.args.silent {
		fmt.Fprintf(os.Stdout, "%s\n", response.body)
	}

	return nil
}

func (v Validator) makeRequest(url string) (Response, error) {
	headers, err := v.getHeaders()
	if err != nil {
		return Response{}, err
	}

	body, err := v.getBody()
	if err != nil {
		return Response{}, err
	}

	res, err := makeRequest(Request{
		method:  v.args.method,
		url:     url,
		headers: headers,
		body:    body,
	})
	if err != nil {
		return Response{}, v.handleError(fmt.Sprintf(validatorErrorHttpRequestError, err.Error()))
	}

	return res, nil
}

func (v Validator) validateResponse(res Response) error {
	if uint(res.statusCode) != v.args.code {
		return v.handleError(fmt.Sprintf(validatorErrorStatusCodeMismatch, res.statusCode, v.args.code))
	}

	if v.args.schemaFilename == "" {
		return nil
	}

	schemaFilename, err := filepath.Abs(v.args.schemaFilename)
	if err != nil {
		return v.handleError(validatorErrorSchemaFileError)
	}

	schemaFilename = uriFromPath(schemaFilename)

	loadedSchema := gojsonschema.NewReferenceLoader(schemaFilename)
	loadedBody := gojsonschema.NewStringLoader(res.body)
	result, err := gojsonschema.Validate(loadedSchema, loadedBody)
	if err != nil {
		return v.handleError(fmt.Sprintf(validatorErrorGeneric, err.Error()))
	}

	if !result.Valid() {
		return v.handleError(validatorErrorSchemaMismatch)
	}

	return nil
}

func (v Validator) getHeaders() (map[string]string, error) {
	if v.args.headersFilename == "" {
		return nil, nil
	}

	headersBytes, err := ioutil.ReadFile(v.args.headersFilename)
	if err != nil {
		return nil, v.handleError(validatorErrorHeaderFileReadError)
	}

	var headersData map[string]string
	err = json.Unmarshal(headersBytes, &headersData)
	if err != nil {
		return nil, v.handleError(validatorErrorHeaderDeserializeError)
	}

	return headersData, nil
}

func (v Validator) getBody() ([]byte, error) {
	if v.args.bodyFilename == "" {
		return nil, nil
	}

	bodyBytes, err := ioutil.ReadFile(v.args.bodyFilename)
	if err != nil {
		return nil, v.handleError(validatorErrorBodyFileReadError)
	}

	return bodyBytes, nil
}

func (v Validator) handleError(message string) error {
	if v.args.silent {
		return errors.New("")
	}

	return errors.New(message)
}

func uriFromPath(path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "file://") {
		return path
	}

	return "file://" + path
}
