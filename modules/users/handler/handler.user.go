package handler

import (
	"net/http"
	"projecttemplatebyfrans/modules/users"
	"projecttemplatebyfrans/schemas"
	"projecttemplatebyfrans/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService users.Service
}

func InitUserHandler(g *gin.Engine, userService users.Service) {
	handler := &UserHandler{
		UserService: userService,
	}

	routeAPI := g.Group("/api/v1/user")
	routeAPI.POST("/create", handler.CreateUserHandler)
	routeAPI.GET("/get-all", handler.GetAllUsersHandler)
}

// Create User
// @Tags Users
// @Summary Create User
// @Description Create User
// @ID User-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateUserRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/user/create [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	//
	var req schemas.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.UserService.CreateUserService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create User", map[string]interface{}{
		"id": ID,
	})

}

// Get User
// @Tags Users
// @Summary Get User
// @Description Get User
// @ID User-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200  {object} schemas.Response
// @Router /v1/user/get-all [get]
func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.UserService.GetUsersService()
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Users", users)
}
