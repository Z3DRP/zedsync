package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Z3DRP/zedsync/internal/api"
	"github.com/Z3DRP/zedsync/internal/config"
	"github.com/Z3DRP/zedsync/internal/container"
	"github.com/Z3DRP/zedsync/internal/crane"
	"github.com/Z3DRP/zedsync/internal/database/store"
)

func main() {
	fmt.Println("running on 8090")
	services := []string{"user"}
	cfg, err := config.Load()
	if err != nil {
		log.Println("failed to load database config exiting...")
		return
	}

	conn, err := store.DBCon(cfg.Database)
	if err != nil {
		log.Printf("an erro occurred while connecting to db %v", err)
	}
	if err := run(cfg.Server, conn, services); err != nil {
		log.Printf("an error occurred while running server %v", err)
	}
}

func run(serverCfg config.ServerCfg, con *sql.DB, services []string) error {
	logger := crane.DefaultLogger
	dbStore := store.NewBuilder().SetDB(con).SetBunDB().Build()
	// userRepo := urepo.New(dbStore)
	// usrService := usr.New(userRepo)
	// userApi := usr.Initialize(usrService, logger)

	container := container.New(dbStore, logger)
	if err := container.RegisterServices(services); err != nil {
		logger.MustDebugErr(err)
		return err
	}
	server := api.NewServer(serverCfg, container.GetEndpoints())

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.MustDebugErr(err)
		return err
	}
	return nil
}
