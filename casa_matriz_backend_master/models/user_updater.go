package models

type UserUpdater struct {
	Us_id        string `form:"id"`
	Tu_id        string `form:"user-type"`
	Us_nombre    string `form:"name"`
	Us_fechanac  string `form:"birthday"`
	Us_correo    string `form:"email"`
	Us_psswrd    string `form:"password"`
	Us_direccion string `form:"address"`
	Us_telefono  string `form:"phone"`
	Us_activo    string `form:"state"`
}
