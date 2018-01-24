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
	"os"

	"gopkg.in/urfave/cli.v1"
)

// Version contains the current application version
const Version = "0.5.0"

func main() {
	app := cli.NewApp()
	args := setupCli(app, Version)
	app.Action = func(c *cli.Context) error {
		validator := Validator{args: args}
		err := validator.Run(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	}

	app.Run(os.Args)
}
