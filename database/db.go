package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4/database/postgres"

	// source/file import is required for migration files to read
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	// load pq as database driver
	_ "github.com/lib/pq"
)

var (
	Todo *sqlx.DB
)

type SSLMode string

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

// ConnectAndMigrate function connects with a given database and returns error if there is any error
func ConnectAndMigrate(host, port, databaseName, user, password string, sslMode SSLMode) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	Todo = DB
	return migrateUp(DB)
}

func ShutdownDatabase() error {
	return Todo.Close()
}

// migrateUp function migrate the database and handles the migration logic
func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// SetupBindVars prepares the SQL statement for batch insert
func SetupBindVars(stmt, bindVars string, length int) string {
	bindVars += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(bindVars, length))
	return replaceSQL(strings.TrimSuffix(stmt, ","), "?")
}

// replaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func replaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
