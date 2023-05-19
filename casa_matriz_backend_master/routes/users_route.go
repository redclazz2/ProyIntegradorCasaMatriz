package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UsersRoute(c echo.Context) error {
	//Parse BodyRequest
	page, _ := strconv.Atoi(c.Request().Header.Get("page"))
	offset := (page - 1) * 10

	//Get Database
	var db = db.GetDatabase()

	response, err := db.Query("SELECT * FROM usuarios LIMIT 10 OFFSET ?", offset)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	list := []models.UserResponse{}

	for response.Next() {
		user := models.User{}
		if err := response.Scan(&user.Us_id, &user.Tu_id, &user.Us_nombre, &user.Us_fechanac, &user.Us_correo, &user.Us_psswrd, &user.Us_direccion, &user.Us_telefono, &user.Us_documento, &user.Us_foto, &user.Us_activo); err != nil {
			panic(err)
		}
		res := models.UserResponse{Us_id: user.Us_id, Tu_id: user.Tu_id, Us_nombre: user.Us_nombre, Us_fechanac: user.Us_fechanac, Us_correo: user.Us_correo, Us_direccion: user.Us_direccion, Us_telefono: user.Us_telefono, Us_documento: user.Us_documento, Us_foto: user.Us_foto, Us_activo: user.Us_activo}
		list = append(list, res)
	}

	return c.JSON(http.StatusOK, list)

}
