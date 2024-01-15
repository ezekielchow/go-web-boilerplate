package main

import (
	"fmt"
	"go-boilerplate/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func main() {

	env, err := utils.LoadEnv()

	if err != nil {
		panic("env missing")
	}

	utils.SetupDatabase(env.DSN)
	utils.AddValidators()

	r := gin.Default()
	r.Use(cors.Default())

	ApiHandler(r)

	// Form errors
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
		opt.ValidatePrivateFields = true
	})

	log.Fatal(http.ListenAndServe(":"+env.APP_PORT, r))
	fmt.Println("Serving at " + env.APP_PORT)
}
