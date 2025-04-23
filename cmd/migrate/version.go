package migrate

import (
	"log"

	"github.com/spf13/cobra"
)

func migrateVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print current migration version",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := newMigrator()
			if err != nil {
				log.Fatalf("❌ migrate.New error: %v", err)
			}
			v, dirty, err := m.Version()
			if err != nil {
				log.Fatalf("❌ Version error: %v", err)
			}
			log.Printf("📌 Current version: %d (dirty: %v)", v, dirty)
		},
	}
}
