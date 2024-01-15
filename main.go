package main

import (
	"fmt"
	"go-boilerplate/utils"
	"log"
	"net/http"

	"github.com/gookit/validate"
	"github.com/rs/cors"
)

func main() {

	env, err := utils.LoadEnv()

	if err != nil {
		panic("env missing la")
	}

	utils.SetupDatabase(env.DSN)

	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)

	StaticHandler(mux)
	ApiHandler(mux)
	WebHandler(mux)

	// Form errors
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
		opt.ValidatePrivateFields = true
	})

	log.Fatal(http.ListenAndServe(":"+env.APP_PORT, handler))
	fmt.Println("Serving at " + env.APP_PORT)
}
