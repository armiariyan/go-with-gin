package dto

//BorrowerUpdateDTO is used by client when PUT update borrower data
type BorrowerUpdateDTO struct {
	ID         	int64  	  `json:"id" form:"id"`
	Id_borrower string    `json:"id_borrower" form:"id_borrower" binding:""`
	Id_user     int64     `json:"user" form:"id_user" binding:""`
	House       int64	  `json:"house" form:"id_house" binding:"required,oneof=1 2 3"`
}

//BorrowerCreateDTO is is a model that client use to create borrower data
type BorrowerCreateDTO struct {
	Id_borrower string    `json:"id_borrower" form:"id_borrower" binding:""`
	Id_user     int64     `json:"id_user" form:"id_user" binding:""`
	House       int64	  `json:"house" form:"id_house" binding:"required,oneof=1 2 3"`
}
