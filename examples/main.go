package main

import (
	"time"

	"github.com/philchia/gol"
)

func main() {
	gol.Debug("Hello world")
	gol.Info("Hello world")
	gol.Warn("Hello world")
	gol.Error("Hello world")
	gol.Critical("Hello world")
	time.Sleep(time.Second)
}
