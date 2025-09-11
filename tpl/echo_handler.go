package tpl

func EchoHandlerTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package handler

import (
	"test/interfaces"
	"test/service"
)

type NginxgenVncHandler struct {
	nginxgenVncService *service.YourService
}

func NewNginxgenVncHandler(vncSv *service.YourService) interfaces.YourServiceHandler {
	return &NginxgenVncHandler{vncSv}
}

func (h *NginxgenVncHandler) Generate() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}

func (h *NginxgenVncHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}
func (h *NginxgenVncHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}


`)
}
