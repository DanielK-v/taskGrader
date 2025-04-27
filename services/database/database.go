package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DB_USER")
    cfg.Passwd = os.Getenv("DB_PASSWORD")
    cfg.Net = "tcp"
    cfg.Addr = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
    cfg.DBName = os.Getenv("DB_NAME")
	
	var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
