package handlers

import (
	"net/http"
	"fmt"
	authdto "pengirimanbarang/dto/auth"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"
	"pengirimanbarang/pkg/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

type dataRegister struct {
	User interface{} `json:"user"`
}

type dataLogin struct {
	User interface{} `json:"user"`
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

// Register
func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	// Log the request to check its content
	fmt.Printf("Registration Request: %+v\n", request)

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	user := models.User{
		UserName: request.UserName,
		Name:     request.Name,
		Password: password,
	}

	// Log the user object to check its content
	fmt.Printf("User Object: %+v\n", user)

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: dataRegister{User: data}})
}



// Login
func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	user, err := h.AuthRepository.Login(request.UserName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	// Check if the user is active
	if user.Status != "active" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "User is not active"})
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Wrong email or password"})
	}

	loginResponse := authdto.LoginResponse{
		ID: user.ID,
		UserName: user.UserName,
		Name:     user.Name,
		Password: user.Password,
		Role: user.Role,
		Status: user.Status,
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: dataLogin{User: loginResponse}})
}



// CheckAuth
func (h *handlerAuth) CheckAuth(c echo.Context) error {
	// Retrieve user ID from the context (assuming it's stored in the session or another mechanism)
	userID := getUserIdFromContext(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "failed", Message: "Unauthorized"})
	}

	// Use the user ID to fetch user information from the repository
	user, err := h.AuthRepository.CheckAuth(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "failed", Message: "Unauthorized"})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: user})
}

// Assuming you have a function to get the user ID from the context
func getUserIdFromContext(c echo.Context) int {

	return 0
}
