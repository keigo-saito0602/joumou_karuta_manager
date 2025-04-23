package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-delve/delve/pkg/config"
	"github.com/keigo-saito0602/joumou_karuta_manager/config/logger"
	"github.com/keigo-saito0602/joumou_karuta_manager/di"
	"github.com/keigo-saito0602/joumou_karuta_manager/router"

	"github.com/labstack/echo/v4"
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
	config.LoadConfig()
	e := echo.New()

	container := di.NewContainer()
	log.Println("✅ Connected to MySQL:", container.DB)

	router.RegisterRoutes(e, container.Handlers)
	logger.Init()

	go func() {
		log.Println("🚀 Server is running at http://localhost:8080/swagger/index.html")
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("👋 Shutting down gracefully... Bye!")
	return nil
}
