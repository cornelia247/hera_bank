package main

import (
	"database/sql"
	"log"

	"github.com/cornelia247/hera_bank/api"
	db "github.com/cornelia247/hera_bank/db/sqlc"
	"github.com/cornelia247/hera_bank/util"
	_ "github.com/lib/pq"
	_ "go.uber.org/mock/mockgen/model"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
