package v1

import (
	"go-clean/internal/entity"
	"go-clean/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests for user management
type UserController struct {
	userUseCase usecase.UserUseCase
}

// NewUserController creates a new UserController with the given userUseCase
func newUserController(handler *gin.RouterGroup, userUseCase usecase.UserUseCase) *UserController {
	uc := &UserController{
		userUseCase: userUseCase,
	}
	handler.POST("/register", uc.Register)
	handler.POST("/login", uc.Login)
	return uc
}

// RegisterReq defines the structure for the user registration request.
type RegisterReq struct {
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

// Register handles the user registration process
func (uc *UserController) Register(c *gin.Context) {
	var request RegisterReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Transform DTO to entity for use case logic. This layer can also hash the password.
	newUser := entity.User{
		Email:    request.Email,
		Password: request.Password, // Consider hashing the password before saving
		Name:     request.Name,
	}

	err := uc.userUseCase.Register(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully."})
}

var loginInfoReq struct {
	Mobile   string `json:"mobile" `
	Password string `json:"password"`
}

// Login handles the user login process
func (uc *UserController) Login(c *gin.Context) {
	if err := c.ShouldBindJSON(&loginInfoReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUseCase.Login(c.Request.Context(), loginInfoReq.Mobile, loginInfoReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Assuming creating and returning JWT token or similar logic to be implemented here
	token := "dummy_jwt_token_for_" + user.Email // Replace with actual token generation logic

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
