package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database instance
var DB *sqlx.DB

// Connect function
func Connect() error {
	var err error
	p := os.Getenv("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	DB, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Connection Opened to Database ✅")

	return nil
}
