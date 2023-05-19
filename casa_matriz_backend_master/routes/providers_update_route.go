package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func UpdateProviderRoute(c echo.Context) error {
	// Obtener los datos del usuario de la solicitud HTTP
	provider := new(models.ProvidersUpdater)
	if err := c.Bind(provider); err != nil {
		return err
	}

	//Validar presencia de un id
	if provider.Pro_id == "" {
		return c.String(http.StatusBadRequest, "Debes proporcionar el id del usuario a actualizar")
	}

	// Preparar la sentencia SQL de actualización dinámica
	var setValues []string
	var args []interface{}
	query := "UPDATE proveedores SET "
	//Nombre
	if provider.Pro_nombre != "" {
		setValues = append(setValues, "pro_nombre = ?")
		args = append(args, provider.Pro_nombre)
	}
	//Nit
	if provider.Pro_nit != "" {
		setValues = append(setValues, "pro_nit = ?")
		args = append(args, provider.Pro_nit)
	}
	//Rut
	if provider.Pro_rut != "" {
		setValues = append(setValues, "pro_rut = ?")
		args = append(args, provider.Pro_rut)
	}
	//Country
	if provider.Pro_pais != "" {
		setValues = append(setValues, "pro_pais = ?")
		args = append(args, provider.Pro_pais)
	}
	//Activo
	if provider.Pro_activo != "" {
		setValues = append(setValues, "pro_activo = ?")
		val, err := strconv.ParseBool(provider.Pro_activo)
		if err != nil {
			return err
		}
		args = append(args, val)
	}

	// Unir los valores de actualización en la sentencia SQL
	query += strings.Join(setValues, ", ")
	query += " WHERE pro_id = ?"
	args = append(args, provider.Pro_id)

	// Ejecutar la sentencia SQL de actualización
	var db = db.GetDatabase()
	result, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Comprobar si se actualizó el usuario
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return c.String(http.StatusBadRequest, "No se pudo actualizar el proveedor / no habían cambios que realizar")
	}
	return c.String(http.StatusOK, "Proveedor actualizado exitosamente")
}
