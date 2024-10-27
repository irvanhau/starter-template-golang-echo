package handler

type RegisterInput struct {
	Username    string `json:"username" form:"username" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type LoginInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
