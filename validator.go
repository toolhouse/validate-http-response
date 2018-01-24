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

const MISSING_URL = "URL is required."
const HEADER_FILE_READ_ERROR = "Could not read the provided headers file."
const HEADER_DESERIALIZE_ERROR = "Could not deserialize headers."
const BODY_FILE_READ_ERROR = "Could not read the provided body file."
const HTTP_REQUEST_ERROR = "HTTP error: \"%s\"."
const STATUS_CODE_MISMATCH = "Actual status code (%d) does not match expected status code (%d)."
const SCHEMA_FILE_ERROR = "Could not use provided schema file name."
const VALIDATION_ERROR = "Validation error: \"%s\"."
const SCHEMA_MISMATCH = "Response did not match the provided schema."

type Validator struct {
	args *Args
}

func (v Validator) Run(url string) error {
	if url == "" {
		return v.HandleError(MISSING_URL)
	}

	response, err := v.MakeRequest(url)
	if err != nil {
		return err
	}

	err = v.ValidateResponse(response)
	if err != nil {
		return err
	}

	if !v.args.silent {
		fmt.Fprintf(os.Stdout, "%s\n", response.body)
	}

	return nil
}

func (v Validator) MakeRequest(url string) (Response, error) {
	headers, err := v.GetHeaders()
	if err != nil {
		return Response{}, err
	}

	body, err := v.GetBody()
	if err != nil {
		return Response{}, err
	}

	res, err := makeRequest(Request{
		method: v.args.method,
		url: url,
		headers: headers,
		body: body,
	})
	if err != nil {
		return Response{}, v.HandleError(fmt.Sprintf(HTTP_REQUEST_ERROR, err.Error()))
	}

	return res, nil
}

func (v Validator) ValidateResponse(res Response) error {
	if uint(res.statusCode) != v.args.code {
		return v.HandleError(fmt.Sprintf(STATUS_CODE_MISMATCH, res.statusCode, v.args.code))
	}

	if v.args.schemaFilename == "" {
		return nil
	}

	schemaFilename, err := filepath.Abs(v.args.schemaFilename)
	if err != nil {
		return v.HandleError(SCHEMA_FILE_ERROR)
	}

	schemaFilename = URIFromPath(schemaFilename)

	loadedSchema := gojsonschema.NewReferenceLoader(schemaFilename)
	loadedBody := gojsonschema.NewStringLoader(res.body)
	result, err := gojsonschema.Validate(loadedSchema, loadedBody)
	if err != nil {
		return v.HandleError(fmt.Sprintf(VALIDATION_ERROR, err.Error()))
	}

	if !result.Valid() {
		return v.HandleError(SCHEMA_MISMATCH)
	}

	return nil
}

func (v Validator) GetHeaders() (map[string]string, error) {
	if v.args.headersFilename == "" {
		return nil, nil
	}

	headersBytes, err := ioutil.ReadFile(v.args.headersFilename)
	if err != nil {
		return nil, v.HandleError(HEADER_FILE_READ_ERROR)
	}

	var headersData map[string]string
	err = json.Unmarshal(headersBytes, &headersData)
	if err != nil {
		return nil, v.HandleError(HEADER_DESERIALIZE_ERROR)
	}

	return headersData, nil
}

func (v Validator) GetBody() ([]byte, error) {
	if v.args.bodyFilename == "" {
		return nil, nil
	}

	bodyBytes, err := ioutil.ReadFile(v.args.bodyFilename)
	if err != nil {
		return nil, v.HandleError(BODY_FILE_READ_ERROR)
	}

	return bodyBytes, nil
}

func (v Validator) HandleError(message string) error {
	if v.args.silent {
		return errors.New("")
	}

	return errors.New(message)
}

func URIFromPath(path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "file://") {
		return path
	}

	return "file://" + path
}
