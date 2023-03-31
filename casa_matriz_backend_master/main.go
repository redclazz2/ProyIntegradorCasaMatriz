package main

import (
	"casa_matriz_backend/db"
	router "casa_matriz_backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//Initiate webServer
	e := echo.New()

	//Database Conn
	db.TestDB()

	//Set Routes
	router.ConfigRoutes(e)

	//Serve website
	e.Use(middleware.Static("./static"))

	//Start webserver
	e.Logger.Fatal(e.Start("25.2.104.59:443"))
}
