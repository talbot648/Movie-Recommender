package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gopkg.in/ini.v1"
)

var DB *sql.DB

func InitDB(connectionString string) error {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	DB = db

	// Check if the connection is established
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}
	fmt.Println("Successfully connected to the Database!")
	return nil
}

func GetDBConnectionString(path string) (string, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return "", fmt.Errorf("failed to load config file: %w", err)
	}
	section := cfg.Section("postgres")
	host := section.Key("host").String()
	port := section.Key("port").String()
	user := section.Key("username").String()
	password := section.Key("password").String()
	database := section.Key("db").String()
	sslmode := section.Key("sslmode").String()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, database, sslmode)

	return connectionString, nil
}
