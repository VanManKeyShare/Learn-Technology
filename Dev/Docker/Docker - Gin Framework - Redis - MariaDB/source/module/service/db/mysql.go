package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init_MySQL_DB() error {

	// LẤY THÔNG SỐ TỪ ENVIRONMENT
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// VALIDATE ENV VARS
	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		return fmt.Errorf("MISSING REQUIRED DATABASE ENVIRONMENT VARIABLES")
	}

	// XÂY DỰNG DSN (DATA SOURCE NAME)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// KẾT NỐI TỚI DATABASE
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("FAILED TO OPEN DATABASE: %w", err)
	}

	// CẤU HÌNH CONNECTION POOL
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	// KIỂM TRA KẾT NỐI
	if err = db.Ping(); err != nil {
		return fmt.Errorf("FAILED TO PING DATABASE: %w", err)
	}

	DB = db
	fmt.Println("DATABASE CONNECTED:", dbHost, dbPort)
	return nil
}
