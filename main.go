package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jrpikong/simplebank/api"
	db "github.com/jrpikong/simplebank/db/sqlc"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	ServerAddress = "0.0.0.0:8001"
)

func main() {
	// Logging to a file.
	f, _ := os.Create("storage/simple_bank.log")
	gin.DefaultWriter = io.MultiWriter(f)
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
