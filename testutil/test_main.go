package testutil

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/joumou_karuta_manager")
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()

	seedSQL, err := os.ReadFile("assets/testdata/test_seed.sql")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := testDB.Exec(string(seedSQL)); err != nil {
		log.Fatalf("failed to execute seed SQL: %v", err)
	}

	os.Exit(m.Run())
}
