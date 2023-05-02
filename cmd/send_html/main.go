package main

import (
	"log"
	"net/http"

	"github.com/piliphulko/marketplace-example/internal/pkg/httpout"
)

func main() {
	log.Fatal(
		http.ListenAndServe(":8080", httpout.RouterHTML()),
	)
}
