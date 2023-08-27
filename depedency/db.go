package depedency

import (
	"context"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnectionPool() *pgxpool.Pool {
	url := "user=postgres dbname=postgres password=admin sslmode=disable"
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic("error parsing config" + err.Error())
	}
	// add decimal support
	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxdecimal.Register(c.TypeMap())
		return nil
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("cant connect to database" + err.Error())
	}

	return dbPool
}
