package main

import (
	"log"
	"net/http"

	"github.com/IndraGunawan/gosample"
	"github.com/IndraGunawan/gosample/database"
	"github.com/IndraGunawan/gosample/env"
	"github.com/IndraGunawan/gosample/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	databaseOpt := database.Option{
		Host:     env.GetWithDefault("MYSQL_HOST", "127.0.0.1"),
		Port:     env.GetWithDefault("MYSQL_PORT", "3306"),
		User:     env.Get("MYSQL_USER"),
		Password: env.Get("MYSQL_PASSWORD"),
		Database: env.Get("MYSQL_DATABASE"),
		Charset:  env.GetWithDefault("MYSQL_CHARSET", "utf8"),
	}

	mysql, _ := database.New(databaseOpt)
	userRepository := database.NewUserRepository(mysql)
	userService := gosample.NewUserService(userRepository)

	healthzHandler := handler.NewHealthzHandler()
	userHandler := handler.NewUserHandler(userService)

	router.GET("/healthz", healthzHandler.Healthz)
	router.GET("/users", userHandler.GetAll)
	router.POST("/users", userHandler.Create)
	router.GET("/users/:id", userHandler.GetByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
