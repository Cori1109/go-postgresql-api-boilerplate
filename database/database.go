package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/config"
)

// Database instance
var DB *sql.DB

// Connect function
func Connect() error {
	var err error
	p := config.Dot("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Dot("DB_HOST"), port, config.Dot("DB_USER"), config.Dot("DB_PASSWORD"), config.Dot("DB_NAME")))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Connection Opened to Database ✅")

	return nil
}
