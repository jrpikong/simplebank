package main

import (
	"database/sql"
	"github.com/jrpikong/simplebank/api"
	db "github.com/jrpikong/simplebank/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	ServerAddress = "0.0.0.0:8001"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
