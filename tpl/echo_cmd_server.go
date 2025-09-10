package tpl

func EchoCmdServerTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"net/http"
)

type ApiServer struct {
	systemSecret string
	port         string
	timezone     string
	loglevel     string

	groupSystemAPI *echo.Group
	echo           *echo.Echo
}

var (
	serverCmd = &cobra.Command{
		Use:     "server [command name]",
		Aliases: []string{"command"},
		Short:   "Run api server",
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var comps []string
			if len(args) == 0 {
				comps = cobra.AppendActiveHelp(comps, "Please specify the name for the new command")
			} else if len(args) == 1 {
				comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
			} else {
				comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
			}
			return comps, cobra.ShellCompDirectiveNoFileComp
		},
		Run: func(cmd *cobra.Command, args []string) {
			/*			if len(args) < 1 {
							cobra.CheckErr(fmt.Errorf("add needs a name for the command"))
						}
			*/
			global.server.Start()
		},
	}
)

func (server *ApiServer) Start() {
	server.echo = echo.New()

	server.dependenciesInjection()
	server.setMiddleware()
	server.route()

	if err := server.echo.Start(":" + server.port); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		server.echo.Logger.Fatal("shutting down the server")
	}
}

`)
}
