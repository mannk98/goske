package tpl

func EchoCmdServerMiiddlewareTemplate() []byte {
	return []byte(`/*
{{ .Project.GetCopyright }}
*/
package cmd

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mannk98/golibs/response"
	"io"
	"net/http"
	"os"
)

func (server *ApiServer) setMiddleware() {
	server.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	server.echo.Use(middleware.Logger())
	//server.echo.Use(middleware.Recover())
	server.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			os.Setenv("PROCESS_ID", uuid.New().String())
			return next(c)
		}
	})

	server.echo.HTTPErrorHandler = func(err error, c echo.Context) {

		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		resp := response.NewApiResponse(c.Path())
		resp.Status.Code = code
		resp.Status.Type = http.StatusText(code)
		resp.Data = nil
		reqBody, _ := io.ReadAll(c.Request().Body)
		var decodeBody interface{}
		_ = json.Unmarshal(reqBody, &decodeBody)
		resp.Request = decodeBody

		c.JSON(code, resp)
	}
}

func (server *ApiServer) AuthSecretMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			secret := c.Request().Header.Get("secret")
			if secret == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide secret credentials")
			}

			if secret != server.systemSecret {
				return echo.NewHTTPError(http.StatusForbidden, "Wrong secret credentials")
			}

			return next(c)
		}
	}
}

`)
}
