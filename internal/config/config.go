package config

import (
	"log"
	"os"
	"strings"
)

var ConfigDirectory string

func init() {
	s, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ways := strings.Split(s, `\`)
	for i := 0; i != len(ways); i++ {
		if ways[i] == "restapi-postgresql" {
			break
		}
		ConfigDirectory += ways[i] + "/"
	}
	ConfigDirectory += "restapi-postgresql/internal/config/config.json"
}
