package cmd

import (
	"fmt"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var myEnv = config.ReadEnv()
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of App",
	Long:  `Versioning app to track changes properly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(myEnv["APP_NAME"] + ": A Golang web app by " + myEnv["AUTHOR"] + ". version: " + myEnv["VERSION"])
	},
}
