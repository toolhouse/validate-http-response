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
)

func main() {
	c, err := configFromEnv()
	if err != nil {
		exitWithError("Configuration issue", err)
	}

	output, err := checkReadiness(c.URL)
	if err != nil {
		exitWithError("Error fetching deployment manifest", err)
	}

	fmt.Fprintf(os.Stdout, "Deployment verified. Output:\n%s\n", output)
	os.Exit(0)
}

func exitWithError(desc string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s", desc, err.Error())
	os.Exit(1)
}
