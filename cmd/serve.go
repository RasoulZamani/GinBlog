package cmd

import (
	"github.com/RasoulZamani/hiGin/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server and serve web app",
	Long:  `This command used to run server on host and port that declared in the .env file`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Serve()
	},
}
