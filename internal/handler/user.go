package handler

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// GetUser godoc
// @Summary Get User
// @Schemes
// @Description
// @Tags User Module
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetUserResp
// @Router /v1/user [get]
func (h *UserHandler) GetUser(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.Err(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetUser(ctx, userCode)
	if err != nil {
		v1.Err(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.Ok(ctx, user)
}

// UpdateUser godoc
// @Summary Update User
// @Schemes
// @Description
// @Tags User Module
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.UpdateUserReq true "params"
// @Success 200 {object} v1.Response
// @Router /v1/user [put]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	var req = new(v1.UpdateUserReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.Err(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateUser(ctx, userCode, req); err != nil {
		v1.Err(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.Ok(ctx, nil)
}

// DeleteUser godoc
// @Summary Delete User
// @Schemes
// @Description
// @Tags User Module
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.Response
// @Router /v1/user [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.Err(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	if err := h.userService.DeleteUser(ctx, userCode); err != nil {
		v1.Err(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.Ok(ctx, nil)
}
