package cli

import (
	"github.com/changyoungkwon/gxample/internal/database/migrate"
	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/spf13/cobra"
)

var rollback = false

// MigrateCmd migrate based on internal/database/migrate
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate",
	Run: func(cmd *cobra.Command, args []string) {
		if rollback {
			if err := migrate.RollbackLast(); err != nil {
				logging.Logger.Errorf("cannot rollback last migration, %v", err)
				panic(err)
			}
			logging.Logger.Info("reset successfully")
			return
		}
		if err := migrate.Migrate(); err != nil {
			logging.Logger.Errorf("errors during migrations, %v", err)
			panic(err)
		}
		logging.Logger.Info("migrate succesfully")
	},
}

func init() {
	MigrateCmd.Flags().BoolVarP(&rollback, "rollback", "r", false, "rollback last migration")
}
