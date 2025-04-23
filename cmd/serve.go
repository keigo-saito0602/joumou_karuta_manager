package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

func newServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("📦 Loading environment...")
			if err := godotenv.Load(); err != nil {
				log.Println("⚠️ .env not found (skipped)")
			}

			log.Println("🔌 Connecting to DB...")
			dbConn := db.NewMySQLDB()
			log.Println("✅ DB connected:", dbConn)

			log.Println("🚀 Starting HTTP server on :8080...")
			e := echo.New()
			e.GET("/ping", func(c echo.Context) error {
				return c.String(200, "pong")
			})
			e.Logger.Fatal(e.Start(":8080"))
		},
	}
}
