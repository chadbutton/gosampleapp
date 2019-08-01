package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"innovativeobjects/sampleapp/logwrapper"
)

var cfgFile string
var ver string

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "sampleapp",
	Short: "A utiliy that determines the health of a machine to be run as a prerequisite to a software package install",
	Long:  `This is a quick sample application that uses Cobra as a CLI library.`,
	Run: func(cmd *cobra.Command, args []string) {
		Check()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logwrapper.InitLogging()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
}

func initConfig() {

	viper.SetConfigType("yaml")
	viper.SetConfigName(cfgFile)
	viper.AddConfigPath(".")

	configErr := viper.ReadInConfig()

	if configErr == nil {
		log.Info(fmt.Sprintf("Using config file:", viper.ConfigFileUsed()))
	} else {
		panic(fmt.Errorf("Could not find valid config file, error: %s \n", configErr))
	}
}
