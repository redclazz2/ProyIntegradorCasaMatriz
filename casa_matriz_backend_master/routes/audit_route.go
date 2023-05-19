package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AuditRoute(c echo.Context) error {
	//Parse BodyRequest
	page, _ := strconv.Atoi(c.Request().Header.Get("page"))
	offset := (page - 1) * 10

	//Get Database
	var db = db.GetDatabase()

	response, err := db.Query("SELECT * FROM audit_transac LIMIT 10 OFFSET ?", offset)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	list := []models.Audit_transac{}

	for response.Next() {
		res := models.Audit_transac{}
		if err := response.Scan(&res.Autr_id, &res.Autr_usuario, &res.Autr_tabla, &res.Autr_id_insertado, &res.Autr_fecha); err != nil {
			panic(err)
		}
		list = append(list, res)
	}

	return c.JSON(http.StatusOK, list)

}
