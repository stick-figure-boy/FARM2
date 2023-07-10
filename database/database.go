package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
	conn, err := sqlx.Open("mysql", getDsn(false))
	if err != nil {
		panic(fmt.Errorf("sqlx open error: %w", err))
	}

	return conn
}

func NewTestDB() *sqlx.DB {
	conn, err := sqlx.Open("mysql", getDsn(true))
	if err != nil {
		panic(fmt.Errorf("sqlx open error: %w", err))
	}

	return conn
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func getDsn(isTesting bool) string {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	if isTesting {
		dbConfig = DBConfig{
			Host:     os.Getenv("TEST_DB_HOST"),
			User:     os.Getenv("TEST_DB_USER"),
			Password: os.Getenv("TEST_DB_PASSWORD"),
			DBName:   os.Getenv("TEST_DB_NAME"),
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DBName)

	return dsn
}

// const migrationFilePath = "file://database/migrations/"

// func newMigrate(db *sqlx.DB) *migrate.Migrate {
// 	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
// 	if err != nil {
// 		panic(fmt.Errorf("mysql get instance error: %w", err))
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(
// 		migrationFilePath,
// 		"mysql",
// 		driver,
// 	)
// 	if err != nil {
// 		panic(fmt.Errorf("sql migration error: %w", err))
// 	}

// 	return m
// }
