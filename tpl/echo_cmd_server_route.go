package tpl

func EchoCmdServerRouteTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

func (server *ApiServer) route() {
	// System API
	server.groupSystemAPI = server.echo.Group("/api/function")
	server.groupSystemAPI.Use(server.AuthSecretMiddleware())
	server.groupSystemAPI.POST("/yourservice", server.yH.Generate())
}

`)
}
