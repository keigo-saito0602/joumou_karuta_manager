package cmd

import "github.com/spf13/cobra"

func commandGroups() []*cobra.Group {
	return []*cobra.Group{
		{
			ID:    "server",
			Title: "Server Commands",
		},
		{
			ID:    "migrate",
			Title: "Migration Commands",
		},
	}
}
