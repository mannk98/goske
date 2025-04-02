package tpl

func AddCommandTemplate() []byte {
	return []byte(`/*
{{ .Project.GetCopyright }}
*/
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/mannk98/netutils"
	"github.com/mannk98/goutils/utils"

	"github.com/spf13/cobra"
)

// {{ .CmdName }}Cmd represents the {{ .CmdName }} command
var {{ .CmdName }}Cmd = &cobra.Command{
	Use:   "{{ .CmdName }}",
	Short: "A brief description of your command",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.` + "`" + `,
	Run: {{ .CmdName }}Run,
}

func init() {
	{{ .CmdParent }}.AddCommand({{ .CmdName }}Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// {{ .CmdName }}Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// {{ .CmdName }}Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func {{ .CmdName }}Run(cmd *cobra.Command, args []string) {
	Logger.Info("cmd called")
}

`)
}
