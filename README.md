# Danger in Go

This is a Go version of the popular [Danger](https://danger.systems/) tool.

## Installation of command line tool

```shell
go install github.com/luno/danger-go/cmd/danger-go@latest
yarn global add danger
```

Requires [Danger JS](https://danger.systems/js) to run properly.

## Integrate into project

1. Create a new directory to house the *dangerfile.go* file. This repo uses `build/ci`.
2. Add a `dangerfile.go` to the directory with the following contents:
```go
package main

import "github.com/luno/danger-go"

func Run(d *danger.T, pr danger.DSL) {
	d.Message("danger-go is running!", "", 0)
}
```
3. Run the following in the directory:
```shell
go mod init dangerfile
go get github.com/luno/danger-go
go mod tidy
```

## Running danger-go locally

The `danger-go` command line tool supports `local`, `pr`, and `ci` commands. `danger-go` wraps the corresponding `danger` (js) commands, so to get information about flags, run `danger <command> --help`.

## CI integration

### GitHub Actions
See `.github/workflows/main.yml` as a reference.
