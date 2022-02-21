package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DatabaseConnection interface {
	connect() *sql.DB
	CreateConnectionUrl() string
}

type PostgresConnection struct {
	host, user, db, password string
	port                     int
	conn                     *sql.DB
}

func (conn PostgresConnection) connect() *sql.DB {
	open, err := sql.Open("postgres", conn.CreateConnectionUrl())
	if err != nil {
		log.Fatal(err.Error())
	}

	return open
}

func (conn PostgresConnection) CreateConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conn.user, conn.password, conn.host, conn.port, conn.db)
}

func createPostgresConnection() PostgresConnection {
	return PostgresConnection{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		db:       os.Getenv("DB_DATABASE"),
		password: os.Getenv("DB_PASSWORD"),
	}
}
