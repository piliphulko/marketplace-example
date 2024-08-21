package donkeypacking

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type closeConn func()

func GetConnPostrgresql(postgresqlURL string) (*pgxpool.Pool, closeConn, error) {
	var i int = 1
	for {
		fmt.Printf("|POSTGRESQL|:connection attempt: %d\n", i)
		pool, err := pgxpool.New(context.Background(), postgresqlURL)
		if err != nil && i > 4 {
			return nil, nil, err
		} else if err == nil {
			fmt.Println("|POSTGRESQL|:connection completed successfully")
			return pool, func() { pool.Close() }, nil
		}
		time.Sleep(time.Duration(i^2*250) * time.Microsecond)
		i++
	}
}
