package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
)

func Run() int {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found or failed to load (skipped)")
	} else {
		log.Println(".env file loaded successfully")
	}

	dbConn := db.NewMySQLDB()
	log.Println("Connected to MySQL successfully:", dbConn)

	log.Println("🚀  Application started successfully")
	return 0
}
