# gol: a high performance async log kit for golang

[![Golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![Build Status](https://travis-ci.org/philchia/gol.svg?branch=master)](https://travis-ci.org/philchia/gol)
[![Coverage Status](https://coveralls.io/repos/github/philchia/gol/badge.svg?branch=master)](https://coveralls.io/github/philchia/gol?branch=master)
[![codebeat badge](https://codebeat.co/badges/2b773e81-8362-4d21-9c76-792a64d3df11)](https://codebeat.co/projects/github-com-philchia-gol)
[![Go Report Card](https://goreportcard.com/badge/github.com/philchia/gol)](https://goreportcard.com/report/github.com/philchia/gol)
[![GoDoc](https://godoc.org/github.com/philchia/gol?status.svg)](https://godoc.org/github.com/philchia/gol)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)

gol is a high performance async log infrastructure for golang, which include several useful log backend adapters, include file/file rotate/stmp/slack/elasticsearch etc...

## Introduce

### Level

gol support various log levels, you can set the logger's level to disable some lower level output

```go
const (
    ALL LogLevel = iota
    DEBUG
    INFO
    WARN
    ERROR
    CRITICAL
)
```

### Built in adapters

gol has several built in adapters

- Console adapter support write log to stderr, and this is the default adapter
- File adapter support write log to file
- File rotate adapter support write log to rotate files
- Smtp adapter support write log to email
- Slack adapter support write log to given slack channel
- ES adapter support write log to elastic search (**under development**)

### Customize backend adapters

You can create any backend adapter which implement the Adapter interface.

Actually Adapter is a alias of io.Writer

```go
type Adapter interface {
    io.WriteCloser
}
```

### Color

gol also include a colorful output

![Colorful output](./assets/screen.png)

## Usage

### Log to console

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Not log to console

```go
import (
    "github.com/philchia/gol"
    "runtime"
)
gol.RemoveAdapter(gol.CONSOLELOGGER)
```

### Log to file

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.AddLogAdapter("file", file.NewAdapter("/var/log/tmp.log"))
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Rotate log to file

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.AddLogAdapter("rotate file", rotatefile.NewAdapter("./temp.log", 6, rotatefile.KB*1))
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Set level

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.SetLevel(gol.ERROR)
gol.Debug("Hello, gol!!!") // this will not print
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Set options

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.SetOption(gol.Llongfile | gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llevel)
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Add adapters

You can implement you own custom adapters which implement the Adapter interface.

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

defer gol.Flush()
gol.SetOption(gol.Llongfile | gol.Ldate | gol.Ltime | gol.Lmicroseconds | gol.Llevel)
gol.AddLogAdapter("anonymous", a)
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)
```

## Installation

    $go get github.com/philchia/gol

or you can use `go get -u` to update the package

## Documentation

For docs, see [Documentation](http://godoc.org/github.com/philchia/gol "GoDoc")  or run:

    $godoc github.com/philchia/gol

## Benchmark

gol include a benchmark against the builtin log package, run `$go test ./... -bench . -benchmem` in your terminal to run the bench

![Benchmark](./assets/bench.png)

## Features

- [X] Log level support
- [X] Customizable log option support
- [X] Async write
- [X] Colorful output
- [X] Flush buffered log
- [X] Toggle console adapter
- [X] Logrotate
- [X] Mail adapter
- [X] Slack adapter
- [X] Level support for single adapter
- [ ] Elastic Search adapter for ELK stack
- [ ] 100% coverage
- [ ] Customizable msg buffer size

## License

gol code is published under the MIT license
