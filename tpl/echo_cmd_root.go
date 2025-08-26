package tpl

func EchoCmdRootTemplate() []byte {
	return []byte(`
package cmd

import (
	"errors"
	"github.com/mannk98/golibs/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	"fmt"
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
		// Search config in wd
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(global.cfgFile)
	}

	// Set Default viper
	viper.SetDefault("Server.Location", "Asia/Ho_Chi_Minh")
	viper.SetDefault("Server.Port", "8090")
	viper.SetDefault("App.LogLevel", "DEBUG")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	notFound := &viper.ConfigFileNotFoundError{}
	switch {
	// if err happend when read config file and it is not configFileNotFoundError type
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		// The config file is optional, we shouldn't exit when the config is not found
		break
	default:
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Load EVN form config file
	global.server.systemSecret = viper.GetString("Server.SystemSecret")
	if global.server.systemSecret == "" {
		err := errors.New("Missing Server.SystemSecret in config file: ~/" + global.cfgFile)
		cobra.CheckErr(err)
	}

	global.server.port = viper.GetInt("Server.Port")
	global.server.location = viper.GetString("Server.Location")
	logLevelString := viper.GetString("Server.LogLevel")

	// Setup logging
	switch logLevelString {
	case "DEBUG":
		global.logLevel = log.DebugLevel
	case "INFO":
		global.logLevel = log.InfoLevel
	case "WARNING":
		global.logLevel = log.WarnLevel
	case "ERROR":
		global.logLevel = log.ErrorLevel
	case "CRITICAL":
		global.logLevel = log.PanicLevel
	default:
		global.logLevel = log.DebugLevel
	}
	err = utils.InitLogger(global.logFile, Logger, global.logLevel)
	cobra.CheckErr(err)
}

`)
}
