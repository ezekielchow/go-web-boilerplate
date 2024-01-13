package main

import (
	"fmt"
	"go-boilerplate/utils"
	"log"
	"net/http"
)

func main() {

	env, err := utils.LoadEnv()

	if err != nil {
		panic("env missing")
	}

	utils.SetupDatabase(env.DSN)

	mux := http.NewServeMux()

	StaticHandler(mux)
	ApiHandler(mux)
	WebHandler(mux)

	log.Fatal(http.ListenAndServe(":"+env.APP_PORT, mux))
	fmt.Println("Serving at " + env.APP_PORT)
}
