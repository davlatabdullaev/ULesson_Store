package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"test/api/models"
)

// CreateUser godoc
// @Router       /user [POST]
// @Summary      Creates a new user
// @Description  create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUser false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	createUser := models.CreateUser{}

	if err := c.ShouldBindJSON(&createUser); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	resp, err := h.services.User().Create(context.Background(), createUser)
	if err != nil {
		handleResponse(c, "error while creating user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, resp)
}

// GetUser godoc
// @Router       /user/{id} [GET]
// @Summary      Gets user
// @Description  get user by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetUser(c *gin.Context) {
	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User().GetUser(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while getting user by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, user)
}

// GetUserList godoc
// @Router       /users [GET]
// @Summary      Get user list
// @Description  get user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.UsersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetUserList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	resp, err := h.services.User().GetUsers(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting users", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateUser godoc
// @Router       /user/{id} [PUT]
// @Summary      Update user
// @Description  update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Param        user body models.UpdateUser true "user"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateUser(c *gin.Context) {
	updateUser := models.UpdateUser{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateUser.ID = uid

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.User().Update(context.Background(), updateUser)
	if err != nil {
		handleResponse(c, "error while updating user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// DeleteUser godoc
// @Router       /user/{id} [DELETE]
// @Summary      Delete user
// @Description  delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteUser(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.User().Delete(context.Background(), models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting user by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}

// UpdateUserPassword godoc
// @Router       /user/{id} [PATCH]
// @Summary      Update user password
// @Description  update user password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Param        user body models.UpdateUserPassword true "user"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateUserPassword(c *gin.Context) {
	updateUserPassword := models.UpdateUserPassword{}

	if err := c.ShouldBindJSON(&updateUserPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateUserPassword.ID = uid.String()

	if err = h.services.User().UpdatePassword(context.Background(), updateUserPassword); err != nil {
		handleResponse(c, "error while updating user password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}