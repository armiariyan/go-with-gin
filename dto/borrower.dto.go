package dto

//BorrowerUpdateDTO is used by client when PUT update borrower data
type BorrowerUpdateDTO struct {
	Id_borrower string    `json:"id_borrower" form:"id_borrower" binding:"required"`
	Id_user     int64     `json:"id_user" form:"id_user" binding:"required"`
	House       int64	  `json:"house" form:"id_house" binding:"required"`
}

//BorrowerCreateDTO is is a model that client use to create borrower data
type BorrowerCreateDTO struct {
	Id_borrower string    `json:"id_borrower" form:"id_borrower"`
	Id_user     int64     `json:"id_user" form:"id_user" binding:"required"`
	House       int64	  `json:"id_house" form:"id_house" binding:"required"`
}
