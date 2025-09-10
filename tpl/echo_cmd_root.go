package tpl

func EchoCmdRootTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package cmd

import (
	"errors"
	loglib "github.com/mannk98/golibs/log"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your application",
	Long:  "",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initViper)
	rootCmd.AddCommand(serverCmd)
}

// initConfig reads in config file and ENV variables if set.
func initViper() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if global.cfgFile != "" {
		viper.SetConfigFile(home + "/" + global.cfgFile)
	} else {
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(global.cfgFile)
	}

	// Set Default viper
	viper.SetDefault("Server.Location", "Asia/Ho_Chi_Minh")
	viper.SetDefault("Server.Port", "8080")
	viper.SetDefault("App.LogLevel", "ERROR")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	notFound := &viper.ConfigFileNotFoundError{}
	switch {
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		// The config file is optional, we shouldn't exit when the config is not found
		break
	default:
		log.Infof("Using config file:", viper.ConfigFileUsed())
	}

	// Load EVN form config file
	global.server.systemSecret = viper.GetString("Server.SystemSecret")
	if global.server.systemSecret == "" {
		global.server.systemSecret = os.Getenv("SYSTEM_SECRET")
		if global.server.systemSecret == "" {
			err := errors.New("Missing Server.SystemSecret in config file: ~/" + global.cfgFile)
			cobra.CheckErr(err)
		}
	}

	global.server.port = viper.GetString("Server.Port")
	if global.server.port == "" {
		global.server.port = os.Getenv("PORT")
	}

	global.server.timezone = viper.GetString("Server.TimeZone")
	if global.server.timezone == "" {
		global.server.timezone = os.Getenv("TIMEZONE")
	}

	logLevelString := viper.GetString("Server.LogLevel")
	if logLevelString == "" {
		logLevelString = os.Getenv("LOG_LEVEL")
	}

	// Setup logging
	switch logLevelString {
	case "DEBUG":

		global.logLevel = loglib.DebugLevel
	case "INFO":
		global.logLevel = loglib.InfoLevel
	case "WARNING":
		global.logLevel = loglib.WarnLevel
	case "ERROR":
		global.logLevel = loglib.ErrorLevel
	case "CRITICAL":
		global.logLevel = loglib.PanicLevel
	default:
		global.logLevel = loglib.ErrorLevel
	}

	log = loglib.NewLogger(loglib.WithLogLevel(global.logLevel))
	cobra.CheckErr(err)	
}

`)
}
