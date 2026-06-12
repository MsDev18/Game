package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
	db, err := sql.Open("mysql", "game:game-pass@(localhost:3307)/game?parseTime=true")
	if err != nil {
		panic(fmt.Errorf("cant't open mysql db: %v", err))
	}

	// see important settings section
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{
		db: db,
	}
}
