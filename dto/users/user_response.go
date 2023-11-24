package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" `
	UserName string `json:"username" `
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	NoHp     string `json:"no_hp"`
	Address  string `json:"address"`
}
