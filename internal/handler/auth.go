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
// @Summary 发送验证码
// @Schemes
// @Description
// @Tags 认证模块
// @Accept json
// @Produce json
// @Param request body v1.SendVerifyCodeRequest true "params"
// @Success 200 {object} v1.Response
// @Router /auth/send-code [post]
func (h *AuthHandler) SendVerifyCode(ctx *gin.Context) {
	req := new(v1.SendVerifyCodeRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.SendVerifyCode(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.SendVerifyCode error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// Login godoc
// @Summary 登录注册
// @Schemes
// @Description
// @Tags 认证模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req = new(v1.LoginRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	token, err := h.userService.Login(ctx, ctx.ClientIP(), req)

	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Login error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

// Logout godoc
// @Summary 注销
// @Schemes
// @Description
// @Tags 认证模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.Response
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	err := h.userService.Logout(ctx, userCode)

	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Logout error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
