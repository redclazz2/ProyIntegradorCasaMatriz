package models

type User struct {
	Us_id        string `json:"id"`
	Tu_id        int `json:"user-type"`
	Us_nombre    string `json:"name"`
	Us_fechanac  string `json:"birthday"`
	Us_correo    string `json:"email"`
	Us_psswrd    string `json:"password"`
	Us_direccion string `json:"address"`
	Us_telefono  string `json:"phone"`
	Us_documento string `json:"document"`
	Us_foto      string `json:"photo"`
	Us_activo    bool `json:"state"`
}
