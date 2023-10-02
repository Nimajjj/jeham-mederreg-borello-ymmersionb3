package repo

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const SQL_PROTOCOL = "tcp"
const ALLOW_NATIVE_PASSWORDS = true

var db *sql.DB

func Init() {
	// .env
	if err := godotenv.Load(".env"); err != nil {
		err = fmt.Errorf("error loading .env file: %s", err)
	}

	// DATABASE
	dbConfig := mysql.Config{
		User:                 os.Getenv("DATABASE_USERNAME"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Addr:                 os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"),
		DBName:               os.Getenv("DATABASE_NAME"),
		Net:                  SQL_PROTOCOL,
		AllowNativePasswords: ALLOW_NATIVE_PASSWORDS,
	}

	dba, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		err = fmt.Errorf("error connecting to database: %s", err)
	}

	db = dba
}

func DB() *sql.DB {
	return db
}
