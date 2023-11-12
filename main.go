package main

import (
	"log"

	"github.com/rhodinemma/go-echo-pg/internal/database"
	"github.com/rhodinemma/go-echo-pg/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}

}
