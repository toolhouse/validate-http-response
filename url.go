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
	"net/http"
	"io/ioutil"
)

func makeRequest(url string) (string, int, error) {
	res, err := http.Get(url)

	if res.StatusCode > 199 || res.StatusCode < 300 {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		return string(body), res.StatusCode, err
	}

	return "", res.StatusCode, err
}
