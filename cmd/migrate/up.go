package migrate

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func migrateUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := newMigrator()
			if err != nil {
				log.Fatalf("❌ migrate.New error: %v", err)
			}
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("❌ m.Up error: %v", err)
			}
			log.Println("✅ Migration UP completed.")
		},
	}
}
