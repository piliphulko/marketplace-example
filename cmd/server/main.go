package main

import (
	"fmt"
	"os"

	_ "github.com/piliphulko/restapi-postgresql/internal/pkg/util" // init config file
	"github.com/spf13/viper"
)

func main() {
	fmt.Println(os.UserHomeDir())
	fmt.Println(viper.GetString("POSTGRESQL.DATABASE_URL"))
}
