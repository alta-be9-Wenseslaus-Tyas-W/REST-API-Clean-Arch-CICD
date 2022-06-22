package main

import (
	"restcleanarch/config"
	"restcleanarch/factory"
	_middlewares "restcleanarch/middlewares"
	_routes "restcleanarch/routes"
)

func main() {
	//Init DB Connection
	dbConn := config.InitDB()

	//Init factory
	presenter := factory.InitFactory(dbConn)

	e := _routes.New(presenter)
	_middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
