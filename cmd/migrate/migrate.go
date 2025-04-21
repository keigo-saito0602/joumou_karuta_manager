package migrate

import (
	"github.com/spf13/cobra"
)

const migrateGroupID = "migrate"

// MigrateCommand returns the `migrate` root command
func MigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate",
		Short:   "Database migration commands",
		GroupID: migrateGroupID,
	}

	cmd.AddCommand(
		migrateUpCommand(),
		migrateDownCommand(),
		migrateVersionCommand(),
	)

	return cmd
}
