package main

import (
	"context"
	"log"
	"os"

	"github.com/piliphulko/restapi-postgresql/internal/pkg/sql"

	"github.com/piliphulko/restapi-postgresql/internal/pkg/pgsql"
	_ "github.com/piliphulko/restapi-postgresql/internal/pkg/util" // init config file
)

func main() {
	dbPool, err := pgsql.GetDBPool(context.Background())
	if err != nil {
		os.Exit(1)
	}
	defer dbPool.Close()
	_, err = dbPool.Exec(context.Background(), sql.SchemeDB)
	if err != nil {
		log.Fatal(err)
	}
}
