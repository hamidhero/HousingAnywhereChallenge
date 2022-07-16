package requests

type GetDataBankLocationRequest struct {
	X   string `json:"x" binding:"required"`
	Y   string `json:"y" binding:"required"`
	Z   string `json:"z" binding:"required"`
	Vel string `json:"vel" binding:"required"`
}
