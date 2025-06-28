package main

import (
	"flag"
	"fmt"

	"github.com/condemo/movie-hub/services/rest/api"
)

func main() {
	addr := flag.String("addr", ":5000", "service port")
	flag.Parse()

	api := api.NewApiServer(*addr)
	fmt.Println("Server Running on port: ", *addr)
	api.Run()
}
