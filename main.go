package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mertcanil/simplebank/api"
	db "github.com/mertcanil/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot conntect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
