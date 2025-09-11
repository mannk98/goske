package tpl

func EchoInterfacesTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package interfaces

import "github.com/labstack/echo/v4"

type YourServiceHandler interface {
	Generate() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

`)
}
