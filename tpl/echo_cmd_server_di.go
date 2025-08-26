package tpl

func EchoCmdServerDiTemplate() []byte {
	return []byte(`
package cmd

func (server *ApiServer) dependenciesInjection() {
}

`)
}
