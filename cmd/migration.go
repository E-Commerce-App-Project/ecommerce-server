package cmd

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/appcontext"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Sync database",
	Long:  `Sync database with the latest schema`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		dbInstance, err := appCtx.GetDBInstance()
		if err != nil {
			panic(err)
		}
		// Regiser migration
		database.Migrate(dbInstance)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
