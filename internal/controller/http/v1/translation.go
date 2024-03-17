package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/internal/controller/http/resp"
	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
)

type translationRoutes struct {
	t usecase.Translation
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.Translation) {
	r := &translationRoutes{t}

	h := handler.Group("/translation")
	{
		h.GET("/history", r.history)
	}
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *translationRoutes) history(c *gin.Context) {
	translations, err := r.t.History(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	resp.Success(c, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}
