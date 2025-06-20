package connectdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func Connectdb() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default or env vars")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing one or more required DB environment variables")
	}

	// Chuẩn connection string của go-mssqldb driver
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, dbname)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
