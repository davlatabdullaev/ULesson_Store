package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateIncome godoc
// @Router       /income [POST]
// @Summary      Creates a new income
// @Description  create a new income
// @Tags         income
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncome(c *gin.Context) {
	resp, err := h.services.Income().Create(context.Background())
	if err != nil {
		handleResponse(c, "error while creating income", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, resp)
}

// GetIncomeByID godoc
//