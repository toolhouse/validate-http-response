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
	"bytes"
	"net/http"
	"io/ioutil"
)

type Request struct {
	method string
	url string
	headers map[string]string
	body []byte
}

type Response struct {
	statusCode int
	body string
}

func makeRequest(req Request) (Response, error) {
	client := &http.Client{}
	httpReq, err := http.NewRequest(req.method, req.url, bytes.NewBuffer(req.body))
	if err != nil {
		return Response{}, err
	}

	for key, value := range req.headers {
		httpReq.Header.Set(key, value)
	}

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return Response{}, err
	}
	defer httpRes.Body.Close()

	body, err := ioutil.ReadAll(httpRes.Body)
	return Response{statusCode: httpRes.StatusCode, body: string(body)}, nil
}
