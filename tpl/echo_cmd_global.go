package tpl

func EchoCmdGlobalTemplate() []byte {
	return []byte(`/*
{{ .Project.GetCopyright }}
*/
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
	cfgFile: ".{{ .AppName }}.toml",
	logFile: "{{ .AppName }}.log",
}

var (
	Logger = log.New()
)
`)
}
