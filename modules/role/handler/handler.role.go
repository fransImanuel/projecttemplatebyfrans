package handler

import (
	"net/http"
	"projecttemplatebyfrans/modules/role"
	"projecttemplatebyfrans/schemas"
	"projecttemplatebyfrans/utils"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	RoleService role.Service
}

func InitUserHandler(g *gin.Engine, RoleService role.Service) {
	handler := &RoleHandler{
		RoleService: RoleService,
	}

	routeAPI := g.Group("/api/v1/role")
	routeAPI.POST("/create", handler.CreateRoleHandler)
	// routeAPI.GET("/get-all", handler.CreateUserHandler)
}

// Create Role
// @Tags Roles
// @Summary Create Role
// @Description Create Role
// @ID Role-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateRoleRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/role/create [post]
func (h *RoleHandler) CreateRoleHandler(c *gin.Context) {
	//
	var req schemas.CreateRoleRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.RoleService.CreateRoleService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create Role", map[string]interface{}{
		"id": ID,
	})

}

// // Get User
// // @Tags Users
// // @Summary Get User
// // @Description Get User
// // @ID User-Get
// // @Security ApiKeyAuth
// // @Accept  json
// // @Produce  json
// // @Success 200  {object} schemas.Response
// // @Router /v1/user/get-all [get]
// func (h *RoleHandler) GetAllUsersHandler(c *gin.Context) {
// 	users, err := h.UserService.GetUsersService()
// 	if err != nil {
// 		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
// 		return
// 	}
// 	utils.APIResponse(c, http.StatusOK, "success", "Success Get Users", users)
// }
