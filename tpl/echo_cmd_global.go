package tpl

func EchoCmdGlobalTemplate() []byte {
	return []byte(`
package cmd

import log "github.com/sirupsen/logrus"

type Global struct {
	server *ApiServer

	logLevel log.Level
	cfgFile  string
	logFile  string
}

var global = Global{
	server:  new(ApiServer),
	cfgFile: ".goapi-jwt.toml",
	logFile: ".goapi-jwt.log",
}

var (
	Logger = log.New()
)
`)
}
