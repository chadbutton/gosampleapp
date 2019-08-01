package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"innovativeobjects/sampleapp/logwrapper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the app version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var verStr = fmt.Sprintf(version())
		log.Info(verStr)
	}}

func init() {
	logwrapper.InitLogging()
	rootCmd.AddCommand(versionCmd)
}

func version() string {
	return viper.Get("app.name").(string) + " v1.0"
}
