package models

type ProvidersUpdater struct {
	Pro_id     string `form:"id"`
	Pro_nombre string `form:"nombre"`
	Pro_nit    string `form:"nit"`
	Pro_rut    string `form:"rut"`
	Pro_pais   string `form:"pais"`
	Pro_activo string `form:"state"`
}
