package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginRoute(c echo.Context) error {
	//Parse BodyRequest
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	correo := json_map["email"].(string)
	password := json_map["password"].(string)

	//Get Database
	var db = db.GetDatabase()

	//Injección SQL a la base de datos
	response, err := db.Query("SELECT us_id, tu_id, us_nombre, us_fechanac, us_correo, us_psswrd, us_direccion, us_telefono, us_documento, us_activo FROM usuarios WHERE us_correo = ? AND us_psswrd = ? AND us_activo = 1", correo, password)
	//validación -> Si hay un error finalizar función
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//Mirar si existe el usuario
	user := models.User{}
	if response.Next() {
		response.Scan(&user.Us_id, &user.Tu_id, &user.Us_nombre, &user.Us_fechanac, &user.Us_correo, &user.Us_psswrd, &user.Us_direccion, &user.Us_telefono, &user.Us_documento, &user.Us_activo)

	}

	if user.Us_psswrd == password {

		//Del password
		user.Us_psswrd = "Not Allowed"
		return c.JSON(http.StatusOK, &user)
	}

	return c.String(http.StatusBadRequest, "Invalid Credentials")
}
