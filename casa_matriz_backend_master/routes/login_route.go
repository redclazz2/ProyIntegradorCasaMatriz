package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"casa_matriz_backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginRoute(c echo.Context) error {
	//Parse BodyRequest

	correo := c.Request().Header.Get("email")
	password := utils.Sha512hash(c.Request().Header.Get("password"))

	//Get Database
	var db = db.GetDatabase()

	//Injección SQL a la base de datos
	response, err := db.Query("SELECT us_id, tu_id, us_nombre, us_fechanac, us_correo, us_psswrd, us_direccion, us_telefono, us_documento, us_foto, us_activo FROM usuarios WHERE us_correo = ? AND us_psswrd = ? AND us_activo = 1", correo, password)
	//validación -> Si hay un error finalizar función
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//Mirar si existe el usuario
	user := models.User{}
	if response.Next() {
		response.Scan(&user.Us_id, &user.Tu_id, &user.Us_nombre, &user.Us_fechanac, &user.Us_correo, &user.Us_psswrd, &user.Us_direccion, &user.Us_telefono, &user.Us_documento, &user.Us_foto, &user.Us_activo)

	}

	if user.Us_psswrd == password {

		userResponse := models.UserResponse{Us_id: user.Us_id, Tu_id: user.Tu_id, Us_nombre: user.Us_nombre, Us_fechanac: user.Us_fechanac, Us_correo: user.Us_correo, Us_direccion: user.Us_direccion, Us_telefono: user.Us_telefono, Us_documento: user.Us_documento, Us_foto: user.Us_foto, Us_activo: user.Us_activo}
		return c.JSON(http.StatusOK, &userResponse)
	}

	return c.String(http.StatusBadRequest, "Invalid Credentials")
}
