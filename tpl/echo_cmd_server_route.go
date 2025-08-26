package tpl

func EchoCmdServerRouteTemplate() []byte {
	return []byte(`
package cmd

func (server *ApiServer) route() {
	// System API
	server.groupSystemAPI = server.echo.Group("/api/function")
	server.groupSystemAPI.Use(server.AuthSecretMiddleware())
	/* both call by authentication service when user login */
	//server.groupSystemAPI.POST("/sign", server.jwtHandler.HandleJwtSign())
}

`)
}
