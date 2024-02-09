package handler

import (
	"context"
	"net/http"
	"strconv"
	"test/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBasketProduct godoc
// @Router       /basketProduct [POST]
// @Summary      Creates a new basketProduct
// @Description  create a new basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateBasketProduct false "basket"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasketProduct(c *gin.Context) {
	basketProduct := models.CreateBasketProduct{}

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	createdBasketProduct, err := h.services.BasketProduct().Create(context.Background(), basketProduct)
	if err != nil {
		handleResponse(c, "error is while creating basket product", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, createdBasketProduct)
}

// GetBasketProduct godoc
// @Router       /basketProduct/{id} [GET]
// @Summary      Get basketProduct by id
// @Description  get basketProduct by id
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	resp, err := h.services.BasketProduct().Get(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// GetBasketProductList godoc
// @Router       /basketProducts [GET]
// @Summary      Get basket list
// @Description  get basket list
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Param 	  	 basket_id query string false "basket_id"
// @Success      201  {object}  models.BasketProductResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProductList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
		basketID    string
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	bID := c.Query("basket_id")
	if bID != "" {
		bUID, err := uuid.Parse(bID)
		if err != nil {
			handleResponse(c, "error in basket id", http.StatusBadRequest, err.Error())
			return
		}

		basketID = bUID.String()
	}

	resp, err := h.services.BasketProduct().GetList(context.Background(), models.GetListRequest{
		Page:     page,
		Limit:    limit,
		Search:   search,
		BasketID: basketID,
	})

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateBasketProduct godoc
// @Router       /basketProduct/{id} [PUT]
// @Summary      Update basketProduct
// @Description  update basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Param        basketProduct body models.UpdateBasketProduct false "basketProduct"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasketProduct(c *gin.Context) {
	basketProduct := models.UpdateBasketProduct{}
	uid := c.Param("id")

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	basketProduct.ID = uid

	resp, err := h.services.BasketProduct().Update(context.Background(), basketProduct)
	if err != nil {
		handleResponse(c, "error is while updating basket", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// DeleteBasketProduct godoc
// @Router       /basketProduct/{id} [Delete]
// @Summary      Delete basketProduct
// @Description  delete basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	if err := h.services.BasketProduct().Delete(context.Background(), models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while deleting", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "basket product deleted!")
}
