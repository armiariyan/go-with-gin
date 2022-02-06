package dto

//LenderUpdateDTO is used by client when PUT update lender data
type LenderUpdateDTO struct {
	Id_lender 	string    `json:"id_lender" form:"id_lender" binding:""`
	Id_user     int64     `json:"id_user" form:"id_user" binding:""`
	Sumber_dana string	  `json:"sumber_dana" form:"sumber_dana" binding:"required"`
}

//LenderCreateDTO is is a model that client use to create lender data
type LenderCreateDTO struct {
	Id_lender 	string    `json:"id_lender" form:"id_lender"`
	Id_user     int64     `json:"id_user" form:"id_user" binding:""`
	Sumber_dana string	  `json:"sumber_dana" form:"sumber_dana" binding:"required"`
}
