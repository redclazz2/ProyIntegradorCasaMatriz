package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProvidersRoute(c echo.Context) error {
	//Parse BodyRequest
	page, _ := strconv.Atoi(c.Request().Header.Get("page"))
	offset := (page - 1) * 10

	//Get Database
	var db = db.GetDatabase()

	response, err := db.Query("SELECT * FROM proveedores LIMIT 10 OFFSET ?", offset)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	list := []models.Providers{}

	for response.Next() {
		res := models.Providers{}
		if err := response.Scan(&res.Pro_id, &res.Pro_nombre, &res.Pro_nit, &res.Pro_rut, &res.Pro_pais, &res.Pro_activo); err != nil {
			panic(err)
		}
		list = append(list, res)
	}

	return c.JSON(http.StatusOK, list)

}
