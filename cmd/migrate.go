/*
Copyright © 2025 Ahmed Mohamed <ahmedmohamed24.dev@gmail.com>
*/
package cmd

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/config"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations to update the schema of the database`,
	Run: func(cmd *cobra.Command, args []string) {
		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			log.Errorln(err)
			return
		}
		err = doMigration(step)
		if err != nil {
			log.Errorln(err)
			return
		}
		log.Infoln("Migration completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Int("step", 0, "Migration step")
}
func doMigration(step int) error {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://data/migrations/", "postgres", driver)
	if err != nil {
		return err
	}
	if step != 0 {
		return m.Steps(step)
	}
	return m.Up()
}
