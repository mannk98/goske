package tpl

func EchoCmdServerDiTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

func (server *ApiServer) dependenciesInjection() {
	// init services, handlers, database
}

`)
}
