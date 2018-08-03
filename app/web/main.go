package main

import (
	"log"
	"net/http"

	"github.com/IndraGunawan/gosample"
	"github.com/IndraGunawan/gosample/database"
	"github.com/IndraGunawan/gosample/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	databaseOpt := database.Option{
		Host:     gosample.GetEnvWithDefault("MYSQL_HOST", "127.0.0.1"),
		Port:     gosample.GetEnvWithDefault("MYSQL_PORT", "3306"),
		User:     gosample.GetEnv("MYSQL_USER"),
		Password: gosample.GetEnv("MYSQL_PASSWORD"),
		Database: gosample.GetEnv("MYSQL_DATABASE"),
		Charset:  gosample.GetEnvWithDefault("MYSQL_CHARSET", "utf8"),
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
