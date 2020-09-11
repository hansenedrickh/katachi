package common

import (
	"github.com/hansenedrickh/keisatsu"

	"github.com/hansenedrickh/katachi/config"
)

var Keisatsu keisatsu.Service

func Initialize() {
	config.Load()
	k := config.Keisatsu()
	Keisatsu = k.InitKeisatsu()
	setupLogger()
}
