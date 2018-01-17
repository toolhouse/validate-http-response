# Verify URL

[![Version](https://badge.fury.io/gh/toolhouse%2Fverify-url.svg)](https://github.com/toolhouse/verify-url/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/toolhouse/verify-url)](https://goreportcard.com/report/github.com/toolhouse/verify-url) [![codebeat badge](https://codebeat.co/badges/4c4cc430-53ea-4022-a05a-dd9e34534940)](https://codebeat.co/projects/github-com-toolhouse-verify-url-master) [![](https://images.microbadger.com/badges/image/toolhouse/verify-url.svg)](https://microbadger.com/images/toolhouse/verify-url "Docker Image") [![license](https://img.shields.io/github/license/toolhouse/verify-url.svg)](https://github.com/toolhouse/verify-url/blob/master/LICENSE)

A simple tool to verify that a target URL returns a 200 response.

## Why?

Run as part of a CI pipeline to verify the integrity of the result.

## How?

The application is primarily designed to be run inside a Docker container (although it can also be run as a standalone binary). The application is configured through the following environment variables:

| Environment Variable | Description                   | Example                                   |
|----------------------|-------------------------------|-------------------------------------------|
| `URL`                | The URL of the site to check. | http://www.example.com/health/check/route |

`URL` is required.

### Example

For example, to check for a deployment for example.com which should be on tag `v1.3.1` and commit `b252eb498a0791db07496601ebc7a059dd55cfe9`

```shell
docker run --env URL=https://www.example.com/health/check/route toolhouse/verify-url:latest
```
