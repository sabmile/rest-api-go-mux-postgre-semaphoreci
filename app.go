package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init(user, pass, dbname string) {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)
	var err error
	a.DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {

}
