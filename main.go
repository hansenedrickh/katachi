package main

import (
	"github.com/hansenedrickh/katachi/common"
	"github.com/hansenedrickh/katachi/server"
)

func main() {
	common.Initialize()
	defer common.Keisatsu.WatchPanic()

	server.StartAPIServer()
}
