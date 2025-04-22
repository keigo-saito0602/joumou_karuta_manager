package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/spf13/cobra"
)

func serverStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start application server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(cmd.Context())
		},
	}
}

func runServer(ctx context.Context) error {
	conn := db.NewMySQLDB()
	log.Println("✅ Connected to MySQL:", conn)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	log.Println("🚀 Server is running... Press Ctrl+C to stop.")

	<-stop

	log.Println("👋 Shutting down gracefully... Bye!")

	return nil
}
