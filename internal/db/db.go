package db

import (
	"database/sql"
	"fmt"

	"github.com/RubFV/venky/migrations"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func Start(dsn string) (*sql.DB, error) {

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la base de datos (%w)", err)
	}

	db.SetMaxOpenConns(1)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("error al asignar dialecto a Goose (%w)", err)
	}

	goose.SetBaseFS(migrations.Embed)

	if err := goose.Up(db, "."); err != nil {
		return nil, fmt.Errorf("error aplicando las migraciones (%w)", err)
	}

	return db, nil
}
