package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(c echo.Context) error {
	//Validar Método válido
	if c.Request().Method != "POST" {
		return c.String(http.StatusMethodNotAllowed, "Invalid request method")
	}

	//Body response
	tipo_usuario := c.FormValue("user-type")
	nombre := c.FormValue("name")
	fechanac := c.FormValue("birthday")
	correo := c.FormValue("email")
	password := utils.Sha512hash(c.FormValue("password"))
	address := c.FormValue("address")
	phone := c.FormValue("phone")
	id := c.FormValue("id")
	rawPhoto, err := c.FormFile("photo")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//Save photo
	src, err := rawPhoto.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//Get photo extension
	ext := filepath.Ext(rawPhoto.Filename)

	// Destination of photo
	dst, err := os.Create("public/profile-pictures/" + id + ext)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//Get Database
	var db = db.GetDatabase()

	//Register User
	_, err = db.Exec("INSERT INTO usuarios(us_id, tu_id, us_nombre, us_fechanac, us_correo, us_psswrd, us_direccion, us_telefono, us_documento, us_foto, us_activo) VALUES(NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, TRUE)", tipo_usuario, nombre, fechanac, correo, password, address, phone, id, id+ext)
	//validación -> Si hay un error finalizar función
	if err != nil {
		//Validación -> coreo existente
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
