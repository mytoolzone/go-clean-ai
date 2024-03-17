// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"
	// Swagger docs.
	_ "go-clean/docs"
	"go-clean/internal/usecase"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, u usecase.UserUseCase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newUserController(h, u)
	}
}
