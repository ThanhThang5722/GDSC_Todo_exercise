package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var Once sync.Once

func GetDbInstance() *sql.DB {
	Once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		sqlDriverName := os.Getenv("sqlDriverName")
		dataSource := os.Getenv("DB_DSN")

		db, err = sql.Open(sqlDriverName, dataSource)
		if err != nil {
			log.Fatalf("Cannot connect to DB")
		} else {
			fmt.Printf("Connect success to DB")
		}
	})
	return db
}
