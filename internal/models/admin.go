package models

type Admin struct {
	Username string `json:"name"`
	Password string `json:"-"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}
