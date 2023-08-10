package config

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var once sync.Once

var (
	db *pgxpool.Pool
)

func ConnectDB(cfg Config) {
	dbUrl := cfg.GetDBString()
	dbConfig, err := pgxpool.ParseConfig(dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse database config: %v\n", err)
		os.Exit(1)
	}

	dbConfig.MaxConns = int32(cfg.Database.MaxPoolConns)
	dbConfig.ConnConfig.ConnectTimeout = time.Duration(cfg.Database.MaxPoolConns * int(time.Second))
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	once.Do(func() {
		if db == nil {
			db, err = pgxpool.NewWithConfig(context.Background(), dbConfig)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
				os.Exit(1)
			}
		}
	})
}

func GetDBPool() *pgxpool.Pool {
	return db
}
