package tpl

func EchoCmdServerDiTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

import (
	"test/handler"
	"test/service"
)

func (server *ApiServer) dependenciesInjection() {
	// init services, handlers, database
	server.yS = service.NewYourService("test", 1, nil)
	server.yH = handler.NewNginxgenVncHandler(server.yS)
}

`)
}
