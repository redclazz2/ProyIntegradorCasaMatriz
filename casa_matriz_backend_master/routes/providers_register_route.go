package router

import (
	"casa_matriz_backend/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProvidersRegisterRoute(c echo.Context) error {
	//Validar Método válido
	if c.Request().Method != "POST" {
		return c.String(http.StatusMethodNotAllowed, "Invalid request method")
	}

	//Body response
	nombre := c.FormValue("name")
	nit := c.FormValue("nit")
	rut := c.FormValue("rut")
	pais := c.FormValue("country")

	//Get Database
	var db = db.GetDatabase()

	//Register User
	_, err := db.Exec("INSERT INTO proveedores(pro_id, pro_nombre, pro_nit, pro_rut, pro_pais, pro_activo) VALUES(NULL, ?, ?, ?, ?, TRUE)", nombre, nit, rut, pais)
	//validación -> Si hay un error finalizar función
	if err != nil {
		//Validación -> coreo existente
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
