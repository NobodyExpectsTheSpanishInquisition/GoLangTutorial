package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DatabaseConnection interface {
	connect() *gorm.DB
	getDsn() string
}

type PostgresConnection struct {
	host, user, db, password string
	port                     int
	conn                     *sql.DB
}

func (conn PostgresConnection) connect() *gorm.DB {
	open, err := gorm.Open(postgres.Open(conn.getDsn()), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return open
}

func (conn PostgresConnection) getDsn() string {
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
