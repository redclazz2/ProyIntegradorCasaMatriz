package router

import (
	"casa_matriz_backend/cors"

	"github.com/labstack/echo/v4"
)

func ConfigRoutes(e *echo.Echo) {
	//Test Route
	e.GET("/api/test", TestRoute, cors.CORS())

	//Login Route
	e.GET("/api/login", LoginRoute, cors.CORS())

	//Register Route
	e.POST("/api/register", RegisterRoute, cors.CORS())

	//Audit_transac Route
	e.GET("/api/logs1", AuditRoute, cors.CORS())

	//Users get Route
	e.GET("/api/users", UsersRoute, cors.CORS())

	//Users update Route
	e.PUT("/api/users", UpdateUserRoute, cors.CORS())

	//Providers get Route
	e.GET("/api/providers", ProvidersRoute, cors.CORS())

	//Providers Register Route
	e.POST("/api/providers", ProvidersRegisterRoute, cors.CORS())

	//Providers
	e.PUT("/api/providers", UpdateProviderRoute, cors.CORS())
}
