package router

import (
	"casa_matriz_backend/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(c echo.Context) error {
	//Parse BodyRequest
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//Body response
	tipo_usuario, _ := strconv.Atoi(json_map["user-type"].(string))
	nombre := json_map["name"].(string)
	fechanac := json_map["birthday"].(string)
	correo := json_map["email"].(string)
	password := json_map["password"].(string)
	address := json_map["address"].(string)
	phone := json_map["phone"].(string)
	id := json_map["id"].(string)

	//Get Database
	var db = db.GetDatabase()

	//Register User
	_, err = db.Exec("INSERT INTO usuarios(us_id, tu_id, us_nombre, us_fechanac, us_correo, us_psswrd, us_direccion, us_telefono, us_documento, us_activo) VALUES(NULL, ?, ?, ?, ?, ?, ?, ?, ?, TRUE)", tipo_usuario, nombre, fechanac, correo, password, address, phone, id)
	//validación -> Si hay un error finalizar función
	if err != nil {
		//Validación -> coreo existente
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
