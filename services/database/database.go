package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")

	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	Db.SetMaxOpenConns(25)
	Db.SetMaxIdleConns(5)
	Db.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to the database... ")

	return Db, nil
}
