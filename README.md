# gol is a high performance async log kit for golang

[![Golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![Build Status](https://travis-ci.org/philchia/gol.svg?branch=master)](https://travis-ci.org/philchia/gol)
[![Coverage Status](https://coveralls.io/repos/github/philchia/gol/badge.svg?branch=master)](https://coveralls.io/github/philchia/gol?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/philchia/gol)](https://goreportcard.com/report/github.com/philchia/gol)
[![GoDoc](https://godoc.org/github.com/philchia/gol?status.svg)](https://godoc.org/github.com/philchia/gol)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)

## Introduce

### Async

gol is a high performance async log kit for golang, as of async, gol has a high speed.

### Level

gol support log level

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

### Log rotate

gol has a built in file adapter which support automatically rotate within log files, also customizable rotate count and file size limits

### Customize backend output

You can create any backend adapter which implement the Adapter interface.

Actually Adapter is a alias of io.Writer

```go
type Adapter interface {
    io.Writer
}
```

### Color

gol also include a colorful output.

![Colorful output](./assets/screen.png)


## Usage

### Log to console

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Log to file

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

gol.AddLogAdapter(file.NewFileAdapter("/var/log/tmp.log"))
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)

```

### Set level

```go
import (
    "github.com/philchia/gol"
    "runtime"
)

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

gol.SetOption(gol.Llongfile | gol.Ldate | gol.Ltime | gol.Lmicroseconds)
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

gol.SetOption(gol.Llongfile | gol.Ldate | gol.Ltime | gol.Lmicroseconds)
gol.AddLogAdapter(a)
gol.Debug("Hello, gol!!!")
gol.Criticalf("Hello from %s", runtime.GOOS)
```

## License

gol code is published under MIT license