package cmd

import (
	"github.com/RasoulZamani/hiGin/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database",
	Long:  `This command used to seed database tables with mock data`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Seed()
	},
}

// func migrate() {
// 	bootstrap.Seed()
// }
