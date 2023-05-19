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

	//Redirect HTTP to HTTPs}
	//DESACTIVADO PARA LOCALHOST, el ssl solo sirve si esta toda la infraestructura UP
	//e.Pre(middleware.HTTPSRedirect())

	//Database Conn
	db.TestDB()

	//Set Routes
	router.ConfigRoutes(e)

	//Serve website
	e.Use(middleware.Static("./public"))

	//Start webserver

	//Default
	e.Logger.Fatal(e.Start(":80"))

	// go func(c *echo.Echo) {
	// 	e.Logger.Fatal(e.Start(":80"))
	// }(e)
	// e.Logger.Fatal(e.StartTLS(":443", "./tls/cert.pem", "./tls/cert-key.pem"))
}
