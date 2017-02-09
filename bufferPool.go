package gol

import (
	"sync"

	"bytes"

	"github.com/philchia/gol/level"
)

var bp sync.Pool
var mp sync.Pool

func init() {
	bp.New = func() interface{} {
		return bytes.NewBuffer(nil)
	}
	mp.New = func() interface{} {
		return &logMSG{}
	}
}

type logMSG struct {
	logLevel level.LogLevel
	msg      string
}

func bufferPoolGet() *bytes.Buffer {
	return bp.Get().(*bytes.Buffer)
}

func bufferPoolPut(b *bytes.Buffer) {
	bp.Put(b)
}

func msgPoolGet() *logMSG {
	return mp.Get().(*logMSG)
}

func msgPoolPut(m *logMSG) {
	mp.Put(m)
}
