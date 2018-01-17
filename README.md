# Verify Toolhouse.Monitoring

[![Version](https://badge.fury.io/gh/toolhouse%2Fverify-toolhouse-monitoring.svg)](https://github.com/toolhouse/verify-toolhouse-monitoring/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/toolhouse/verify-toolhouse-monitoring)](https://goreportcard.com/report/github.com/toolhouse/verify-toolhouse-monitoring) [![codebeat badge](https://codebeat.co/badges/4c4cc430-53ea-4022-a05a-dd9e34534940)](https://codebeat.co/projects/github-com-toolhouse-verify-toolhouse-monitoring-master) [![](https://images.microbadger.com/badges/image/toolhouse/verify-toolhouse-monitoring.svg)](https://microbadger.com/images/toolhouse/verify-toolhouse-monitoring "Docker Image") [![license](https://img.shields.io/github/license/toolhouse/verify-toolhouse-monitoring.svg)](https://github.com/toolhouse/verify-toolhouse-monitoring/blob/master/LICENSE)

A simple tool to verify the deployed site/application has properly integrated [Toolhouse.Monitoring](https://github.com/toolhouse/monitoring-dotnet). The application checks the readiness endpoint on the target site.

## Why?

This was developed with two primary use-cases in mind:

- As part of post-deployment tests within a CI/CD pipeline
- To verify/detect the completion of a deployment that includes manual steps

## How?

The application is primarily designed to be run inside a Docker container (although it can also be run as a standalone binary). The application is configured through the following environment variables:

| Environment Variable | Description                   | Example                       |
|----------------------|-------------------------------|-------------------------------|
| `URL`                | The URL of the site to check. | http://www.example.com/health |

`URL` is required.

### Example

For example, to check for a deployment for example.com which should be on tag `v1.3.1` and commit `b252eb498a0791db07496601ebc7a059dd55cfe9`

```shell
docker run --env URL=https://www.example.com/health toolhouse/verify-toolhouse-monitoring:latest
```
