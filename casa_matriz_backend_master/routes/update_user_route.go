package router

import (
	"casa_matriz_backend/db"
	"casa_matriz_backend/models"
	"casa_matriz_backend/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type NewStatus struct {
	New_state string `json:"new-state"`
}

func UpdateUserRoute(c echo.Context) error {
	// Obtener los datos del usuario de la solicitud HTTP
	user := new(models.UserUpdater)
	if err := c.Bind(user); err != nil {
		return err
	}

	//Validar presencia de un id
	if user.Us_id == "" {
		user.Us_id = c.FormValue("id")
		if user.Us_id == "" {
			return c.String(http.StatusBadRequest, "Debes proporcionar el id del usuario a actualizar")
		}
	}

	//En caso de recibir foto, actualizar
	rawPhoto, _ := c.FormFile("photo")
	ext := ""
	if rawPhoto != nil {
		//Save photo
		src, err := rawPhoto.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		//Get photo extension
		ext = filepath.Ext(rawPhoto.Filename)

		// Destination of photo
		dst, err := os.Create("public/profile-pictures/" + user.Us_id + ext)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	// Preparar la sentencia SQL de actualización dinámica
	var setValues []string
	var args []interface{}
	query := "UPDATE usuarios SET "
	//Foto
	if rawPhoto != nil {
		setValues = append(setValues, "us_foto = ?")
		args = append(args, user.Us_id+ext)
	}
	//Tipo de usuario
	if user.Tu_id != "" {
		setValues = append(setValues, "tu_id = ?")
		args = append(args, user.Tu_id)
	}
	//Estado
	if user.Us_activo != "" {
		setValues = append(setValues, "us_activo = ?")
		val, err := strconv.ParseBool(user.Us_activo)
		if err != nil {
			return err
		}
		args = append(args, val)
	}
	//Nombre
	if user.Us_nombre != "" {
		setValues = append(setValues, "us_nombre = ?")
		args = append(args, user.Us_nombre)
	}
	//Fecha Nacimiento
	if user.Us_fechanac != "" {
		setValues = append(setValues, "us_fechanac = ?")
		args = append(args, user.Us_fechanac)
	}
	//Correo
	if user.Us_correo != "" {
		setValues = append(setValues, "us_correo = ?")
		args = append(args, user.Us_correo)
	}
	//Contraseña
	if user.Us_psswrd != "" {
		setValues = append(setValues, "us_psswrd = ?")
		args = append(args, utils.Sha512hash(user.Us_psswrd))
		fmt.Println(args...)
	}
	//Dirección
	if user.Us_direccion != "" {
		setValues = append(setValues, "us_direccion = ?")
		args = append(args, user.Us_direccion)
	}
	//Telefono
	if user.Us_telefono != "" {
		setValues = append(setValues, "us_telefono = ?")
		args = append(args, user.Us_telefono)
	}

	// Comprobar si se recibieron campos para actualizar
	if len(setValues) == 0 {
		return c.String(http.StatusBadRequest, "No se recibieron campos para actualizar")
	}

	// Unir los valores de actualización en la sentencia SQL
	query += strings.Join(setValues, ", ")
	query += " WHERE us_id = ?"
	args = append(args, user.Us_id)

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
		return c.String(http.StatusBadRequest, "No se pudo actualizar el usuario / no habían cambios que realizar")
	}
	return c.String(http.StatusOK, "Usuario actualizado exitosamente")
}
