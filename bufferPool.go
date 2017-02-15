package gol

import (
	"sync"

	"bytes"

	"github.com/philchia/gol/level"
)

var bp sync.Pool
var mp sync.Pool

func init() {

	mp.New = func() interface{} {
		return &logMSG{}
	}
}

type logMSG struct {
	logLevel level.LogLevel
	bf       bytes.Buffer
}

func msgPoolGet() *logMSG {
	return mp.Get().(*logMSG)
}

func msgPoolPut(m *logMSG) {
	m.bf.Reset()
	mp.Put(m)
}
