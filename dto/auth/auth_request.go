package authdto

type AuthRequest struct {
	ID       int    `json:"id"`
	UserName string `json:"username" validate:"required" form:"username"`
	Name    string `json:"name" validate:"required" form:"name"`
	Password string `json:"password" validate:"required" form:"password"`
	Role string `json:"role" form:"role" gorm:"type:varchar(255)"`
	Status   string `json:"status" form:"status" gorm:"type:varchar(255)"`
}

type LoginRequest struct {
	UserName string `json:"username" validate:"required" form:"username"`
	Password string `json:"password" validate:"required" form:"password"`
}


type LoginResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"-"`
	Name string `json:"name"`
	Role string `json:"role"`
	Status   string `json:"status"`
	// Email    string `json:"email"`
	// Token    string `json:"token"`
}
