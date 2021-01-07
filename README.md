[![saurabmish](https://circleci.com/gh/saurabmish/Project-Euler.svg?style=shield)](https://circleci.com/gh/saurabmish/Project-Euler)
[![codecov](https://codecov.io/gh/saurabmish/Project-Euler/branch/master/graph/badge.svg?token=JnMhKH4HfI)](https://codecov.io/gh/saurabmish/Project-Euler)
[![Go Report Card](https://goreportcard.com/badge/github.com/saurabmish/Project-Euler)](https://goreportcard.com/report/github.com/saurabmish/Project-Euler)

# Project Euler

Project Euler problems implemented in Go with TDD.

Since I'm starting with a new programming language, I want to be very confident with the code right from the start. Some of the best practices that I've learned from writing code in Python, including but not limited to - testing, documentation, build environments, etc. will be done in this project from the get go.

<p align="center">
  <img src="project-euler.jpg" />
</p>

## Installation

+ Install OS dependencies

  `xcode-select --install`

+ Install package manager

  `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`

+ Install Golang using the package manager

  `brew install go`

+ Verify installation

  `go version`

## Environment

+ Since this project is created outside `GOPATH`, *Go Modules* have to used. They enable alternative workflows and also solve problems related to dependencies and versioning.

+ In the VCS project root (where `.git` is located), create a module by executing:

  `go mod init`

## Directory and Naming Convention

+ Source code and test are present in the same directory

+ Directories (and everything which doesn't have `.go`) are named in Kebab-Case

+ Filenames have underscores between words

+ Test cases are present in files having `_test` suffix

## Execution

**NOTE**: Commands executed from the project root (where `.git/` and `go.mod` are present

+ ***Program***

  + Execute file

    `go run ./directory-name/file_name.go`

  + Create an executable of the file

    `go build ./directory-name/file_name.go`

+ ***Testing***

  + Run tests associated with the file

    `go test ./directory-name/file_name_test.go`

    **OR *verbosely***

    `go test -v ./directory-name/file_name_test.go`

  + Recursively run *all* tests

    `go test ./...`

+ ***Coverage***

  + Measure code coverage

    + Compute coverage

      `go test -v -coverprofile coverage.out ./directory-name/`

    + Translate raw data to markup language

      `go tool cover -html=coverage.out -o coverage.html`

    + View file (in browser) and analyze info.:

      `open coverage.html`

  + Check total code coverage

    ```
    go test -v -coverprofile total-coverage.out ./...
    go tool cover -html=total-coverage.out -o total-coverage.html
    open total-coverage.html
    ```
