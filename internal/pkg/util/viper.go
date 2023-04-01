package util

import (
	"fmt"
	"log"

	"github.com/piliphulko/restapi-postgresql/internal/config"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	viper.SetConfigFile(config.ConfigDirectory)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	var b bool
	for _, v := range configs {
		if !viper.IsSet(v) {
			fmt.Println("no config: " + v)
			b = true
		}
	}
	if b {
		panic("")
	}
}

var configs []string = []string{
	"POSTGRESQL.DATABASE_URL",
}
