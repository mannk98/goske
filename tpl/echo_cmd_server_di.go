package tpl

func EchoCmdServerDiTemplate() []byte {
	return []byte(`/*
{{ .Project.GetCopyright }}
*/
package cmd

func (server *ApiServer) dependenciesInjection() {
}

`)
}
