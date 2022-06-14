package persistence

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
	Ctx  context.Context
}

// Connect creates a new connection to the database
func Connect(host string, port int, db string, user string, pw string, log *logrus.Logger) (*pgxpool.Pool, error) {
	ctx := context.Background()
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pw, host, port, db)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	//Could move these to envs
	config.MinConns = int32(2)
	config.MaxConns = int32(3)

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	log.Info("Database connection established")

	return pool, nil

}
