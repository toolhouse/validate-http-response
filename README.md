# Verify URL

[![Version](https://badge.fury.io/gh/toolhouse%2Fverify-url.svg)](https://github.com/toolhouse/verify-url/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/toolhouse/verify-url)](https://goreportcard.com/report/github.com/toolhouse/verify-url) [![codebeat badge](https://codebeat.co/badges/4c4cc430-53ea-4022-a05a-dd9e34534940)](https://codebeat.co/projects/github-com-toolhouse-verify-url-master) [![](https://images.microbadger.com/badges/image/toolhouse/verify-url.svg)](https://microbadger.com/images/toolhouse/verify-url "Docker Image") [![license](https://img.shields.io/github/license/toolhouse/verify-url.svg)](https://github.com/toolhouse/verify-url/blob/master/LICENSE)

```
NAME:
   verify-url - Basic testing for URL responses

USAGE:
   verify-url [options] [URL]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --code value, -c value     The expected HTTP status code. Defaults to 200. (default: "200")
   --method value, -m value   The HTTP method to use when calling the URL. Defaults to 'GET'. (default: "GET")
   --schema FILE, --sch FILE  FILE used to load JSON schema for response verification
   --headers FILE, --hd FILE  FILE containing headers to send to URL
   --body FILE, -b FILE       FILE containing body content to send to URL
   --silent, -s               If specified, nothing will be printed to stdout
   --help, -h                 show help
   --version, -v              print the version
```

## Why?

Run as part of a CI pipeline to test integration of changes and check for regressions.

### Example

```shell
./verify-url-darwin_amd64 --schema=validation/baz.json http://www.foo.bar/baz
```

## Building

Install dependencies: `glide install`

Build: `make darwin-amd64` or `make linux-amd64`