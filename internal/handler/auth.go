package handler

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	*Handler
	userService service.UserService
}

func NewAuthHandler(handler *Handler, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		Handler:     handler,
		userService: userService,
	}
}

// SendVerifyCode godoc
// @Summary Send Verification Vode
// @Schemes
// @Description
// @Tags Auth Module
// @Accept json
// @Produce json
// @Param request body v1.SendVerifyCodeReq true "params"
// @Success 200 {object} v1.Response
// @Router /v1/send-code [post]
func (h *AuthHandler) SendVerifyCode(ctx *gin.Context) {
	req := new(v1.SendVerifyCodeReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.Err(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.SendVerifyCode(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.SendVerifyCode error", zap.Error(err))
		v1.Err(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.Ok(ctx, nil)
}

// Login godoc
// @Summary User Registration/Login
// @Schemes
// @Description
// @Tags Auth Module
// @Accept json
// @Produce json
// @Param request body v1.LoginReq true "params"
// @Success 200 {object} v1.Response
// @Router /v1/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req = new(v1.LoginReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.Err(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	token, err := h.userService.Login(ctx, ctx.ClientIP(), req)

	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Login error", zap.Error(err))
		v1.Err(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.Ok(ctx, v1.LoginRespData{
		AccessToken: token,
	})
}

// Logout godoc
// @Summary Logout
// @Schemes
// @Description
// @Tags Auth Module
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.Response
// @Router /v1/logout [post]
func (h *AuthHandler) Logout(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.Err(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	err := h.userService.Logout(ctx, userCode)

	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Logout error", zap.Error(err))
		v1.Err(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.Ok(ctx, nil)
}
