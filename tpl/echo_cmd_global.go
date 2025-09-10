package tpl

func EchoCmdGlobalTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

import (
	loglib "github.com/mannk98/golibs/log"
	"github.com/sirupsen/logrus"
)

type Global struct {
	server *ApiServer

	logLevel logrus.Level
	cfgFile  string
	logFile  string
}

var global = Global{
	server:  new(ApiServer),
	cfgFile: ".{{ .AppName }}.toml",
	logFile: "{{ .AppName }}.log",
}

var (
	log = loglib.NewLogger()
)
`)
}
