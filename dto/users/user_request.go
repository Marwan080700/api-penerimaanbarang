package usersdto

type CreateUserRequest struct {
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	UserName string `json:"username" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required" form:"gender"`
	NoHp     string `json:"no_hp" validate:"required" form:"no_hp"`
	Address  string `json:"address" validate:"required" form:"address"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullName" form:"fullName"`
	UserName string `json:"username" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
