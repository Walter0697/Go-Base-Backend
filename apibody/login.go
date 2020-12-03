package apibody

type LoginCredentials struct {
	Username    string `form:"username"`
	Password 	string `form:"password"`
}

type CreateAccountInput struct {
	Username	string `json:"username" binding:"required"`
	Password	string `json:"password" binding:"required"`
	Userrole	uint   `json:"userrole" binding:"required"`
}

type UpdateAccountInput struct {
	Username	string `json:"username" binding:"required"`
	OldPassword	string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}