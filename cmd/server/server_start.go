package server

import (
	"context"
	"fmt"
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
	// DB接続（例）
	conn := db.NewMySQLDB()
	fmt.Println("✅ Connected to MySQL:", conn)

	// シグナルで終了検知
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🚀 Server is running... Press Ctrl+C to stop.")

	<-stop
	fmt.Println("👋 Shutting down gracefully...")
	return nil
}
