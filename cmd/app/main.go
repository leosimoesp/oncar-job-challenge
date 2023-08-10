package main

import (
	"github.com/leosimoesp/oncar-job-challenge/config"
	"github.com/leosimoesp/oncar-job-challenge/db"
)

func main() {
	cfg := config.Load()
	config.ConnectDB(cfg)

	dbPool := config.GetDBPool()
	defer dbPool.Close()

	args := []string{}
	_ = db.RunMigrate("up", cfg.GetDBString(), "db/migrations", args...)
}
