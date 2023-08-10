package main

import "github.com/leosimoesp/oncar-job-challenge/config"

func main() {
	cfg := config.Load()
	config.ConnectDB(cfg)

	dbPool := config.GetDBPool()
	defer dbPool.Close()
}
