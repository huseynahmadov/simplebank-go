package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
		return
	}

	connPool, _ := pgxpool.New(context.Background(), config.DBSource)
	store := db.NewStore(connPool)

	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
		return
	}
}
