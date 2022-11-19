package main_test

import (
	"log"
	"os"
	"testing"
)

var a main.App

const createTableQuery = ` CREATE TABLE IF NOT EXISTS product
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
	CONSTRAINTS product_key PRIMARY KEY (id)
`

func ensureTableExis() {
	if _, err := a.DB.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM product")
	a.DB.Exec("ALTER SEQUENCE product_id_seq RESTART WITH 1")
}

func TestMain(m *testing.M) {
	a.Init(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	ensureTableExis()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
