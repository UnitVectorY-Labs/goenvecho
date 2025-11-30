[![GitHub release](https://img.shields.io/github/release/UnitVectorY-Labs/goenvecho.svg)](https://github.com/UnitVectorY-Labs/goenvecho/releases/latest) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Active](https://img.shields.io/badge/Status-Active-green)](https://guide.unitvectorylabs.com/bestpractices/status/#active) [![Go Report Card](https://goreportcard.com/badge/github.com/UnitVectorY-Labs/goenvecho)](https://goreportcard.com/report/github.com/UnitVectorY-Labs/goenvecho)

# goenvecho

Simple app that responds to GET requests with a JSON payload listing all environment variables.

## Purpose

This application captures all environment variables that are set and returns them as a JSON object for requests made to `GET /`.

**Why use this?** It’s a straightforward tool for debugging, testing, and development, providing an easy way to view environment variables in a container environment.

**Should I run this in production?** Absolutely not! This could expose sensitive information stored in your environment variables.

## Example Response

Whatever envirionment variables are set will be returned in the JSON, it is that simple.

For exmaple:

```json
{
  "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
  "PORT": "8080",
  "PWD": "/"
}
```

## Configuration

This application runs as a docker container and can utilize the following environment variables:

- `PORT` - The port the application listens on. (default `8080`)

## Get started

Quickly get up and running by pulling the latest release from [goenvecho on GitHub Packages](https://github.com/UnitVectorY-Labs/goenvecho/pkgs/container/goenvecho) and running it in your container environment of choice.
