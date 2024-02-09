package handler

import (
	"context"
	"net/http"
	"test/api/models"

	"github.com/gin-gonic/gin"
)

// CreateIncomeProducts godoc
// @Router       /income_products [POST]
// @Summary      Creates a new income products
// @Description  create a new income products
// @Tags         income_product
// @Accept       json
// @Produce      json
// @Param		 income_products body models.CreateIncomeProducts false "income_products"
// @Success      201  {object}  string
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncomeProducts(c *gin.Context) {
	var incomeProducts = models.CreateIncomeProducts{}

	if err := c.ShouldBindJSON(&incomeProducts); err != nil {
		handleResponse(c, "error while binding json", http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.IncomeProduct().CreateMultiple(context.Background(), incomeProducts)
	if err != nil {
		handleResponse(c, "error while creating incomeProducts", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, "created")
}
