package migrate

import (
	"log"

	"github.com/spf13/cobra"
)

func migrateDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Rollback the last migration",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := newMigrator()
			if err != nil {
				log.Fatalf("❌ migrate.New error: %v", err)
			}
			if err := m.Steps(-1); err != nil {
				log.Fatalf("❌ Migration down error: %v", err)
			}
			log.Println("✅ Migration DOWN executed.")
		},
	}
}
