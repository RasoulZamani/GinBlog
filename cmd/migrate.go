package cmd

import (
	"github.com/RasoulZamani/hiGin/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate models to database tables",
	Long:  `This command used to Migrate models and create database tables`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Migrate()
	},
}

// func migrate() {
// 	bootstrap.Migrate()
// }
