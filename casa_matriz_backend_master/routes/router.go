package router

import "github.com/labstack/echo/v4"

func ConfigRoutes(e *echo.Echo) {
	//Test Route
	e.GET("/api/test", TestRoute)

	//Login Route
	e.GET("/api/login", LoginRoute)

	//Register Route
	e.POST("/api/register", RegisterRoute)
}
