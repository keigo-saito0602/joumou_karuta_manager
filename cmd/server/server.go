package server

import (
	"github.com/spf13/cobra"
)

const serverGroupID = "server"

// ServerCommand returns the `server` root command
func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Short:   "Server operations",
		GroupID: serverGroupID,
	}

	cmd.AddCommand(serverStartCommand())
	return cmd
}
