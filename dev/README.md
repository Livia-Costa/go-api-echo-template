# Dev Manual

<p>The aim of the dev manual is to help developers, providing guidelines, IDE setup, tools, local environment boostrap, helper scripts, and other useful stuff. </p>

##  📔 Table of Contents
<!--ts-->
   * [Pre-requirements](#✂️-pre-requirements)
   * [DevTools](#🔨-devtools)
      * [Instalation](#instalation)
   * [Local environment](#🎨-local-environment)
      * [Start](#start)
      * [Stop](#stop)
   * [Vulnerabilities Check](#🚩-vulnerabilities-check)
   * [Tests](#🔬-tests)
<!--te-->

## ✂️ Pre-requirements
* [Golang 1.20](https://go.dev/doc/install)
* [Docker](https://docs.docker.com/get-docker/)

## 🔨 DevTools

* [golangci-lint](https://golangci-lint.run/)
* [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
* [gomock](https://github.com/golang/mock)
* [go-migrate](https://github.com/golang-migrate/migrate)

### Instalation

To install dev tools, run:

```
./dev/scripts/install-dev-tools.sh
```

## 🎨 Local Environment

### Start
To start local environment, run:

```
./dev/scripts/start-local-env.sh
```

### Stop
To stop local environment, run:

```
./dev/scripts/stop-local-env.sh
```

## 🚩 Vulnerabilities check
To check for vulnerabilities in dependencies, run:

```
make vulncheck
```

## 🔬  Tests
To run all tests, run
```
make test
```