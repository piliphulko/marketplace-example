package pgsql

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func GetDBPool(ctx context.Context) (*pgxpool.Pool, error) {
	var attempt int
	for {
		dbPool, err := pgxpool.New(ctx, viper.GetString("POSTGRESQL.DATABASE_URL"))
		if err != nil {
			if attempt < 5 {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				return nil, err
			}
		}
		return dbPool, nil
	}
}
