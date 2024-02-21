package depedency

import (
	"context"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnectionPool() *pgxpool.Pool {
	url := "user=postgres password=admin host=localhost port=5432 dbname=postgres sslmode=disable"
	// url := "postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10"
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

	if err := dbPool.Ping(context.Background()); err != nil {
		panic("cant ping to db" + err.Error())
	}

	return dbPool
}
