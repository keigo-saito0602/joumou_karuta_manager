package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keigo-saito0602/joumou_karuta_manager/infrastructure/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ .env 読み込み失敗")
	}

	dbConn := db.NewMySQLDB()
	log.Println("✅ DB接続成功！", dbConn)
}
