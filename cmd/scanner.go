package cmd

import (
	"fmt"
	"innovativeobjects/sampleapp/logwrapper"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scannerCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans files for no apparent reason other to demonstrate how quick it is.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var allDirs = []string{"../../../src/github.com", "../../.."}
		Scan(allDirs)
	}}

func init() {
	logwrapper.InitLogging()
	rootCmd.AddCommand(scannerCmd)
}

func Scan(directories []string) {

	log.Info(fmt.Sprintf("filescanner starting..."))

	var ff = func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Info(fmt.Sprintf("error %v at a path %q", err, path))
			return err
		}

		if info.IsDir() {
			log.Info(fmt.Sprintf("filescanner scanning dir: '%v\\%v'", filepath.Dir(path), info.Name()))
		} else {

			var now = time.Now()
			var fileTime = now.Sub(info.ModTime())

			log.Info(fmt.Sprintf("filescanner found file: '%v' modified timestamp: %v filetime %v", info.Name(), info.ModTime(), fileTime))
		}

		return nil
	}

	for _, dir := range directories {

		err := filepath.Walk(dir, ff)

		if err != nil {
			log.Info(fmt.Sprintf("error walking the path %q: %v\n", dir, err))
		}

	}
}
