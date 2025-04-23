package cmd

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/cmd/migrate"
	"github.com/keigo-saito0602/joumou_karuta_manager/cmd/server"
	"github.com/spf13/cobra"
)

func Run() int {
	if err := rootCommand().Execute(); err != nil {
		return 1
	}
	return 0
}

func rootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "joumou",
		Short: "Joumou Karuta Manager CLI",
	}

	for _, g := range commandGroups() {
		root.AddGroup(g)
	}

	root.AddCommand(
		server.ServerCommand(),
		migrate.MigrateCommand(),
		newServeCmd(),
	)

	return root
}
