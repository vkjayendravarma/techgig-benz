package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/vkjayendravarma/techgig-benz-v2/env"
	"github.com/vkjayendravarma/techgig-benz-v2/src/router"
)

func main() {

	r := router.InitRoutes()

	fmt.Println("Server is getting started at http://" + env.Host + ":" + env.Port)
	log.Fatal(http.ListenAndServe(env.Host+":"+env.Port, handlers.LoggingHandler(os.Stdout, r)))
}
