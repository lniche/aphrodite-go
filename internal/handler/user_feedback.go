package handler

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserFeedbackHandler struct {
	*Handler
	userFeedbackService service.UserFeedbackService
}

func NewUserFeedbackHandler(
	handler *Handler,
	userFeedbackService service.UserFeedbackService,
) *UserFeedbackHandler {
	return &UserFeedbackHandler{
		Handler:             handler,
		userFeedbackService: userFeedbackService,
	}
}

// CreateUserFeedback godoc
// @Summary 新增用户反馈
// @Schemes
// @Description
// @Tags 用户反馈模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.CreateUserFeedbackRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user/feedback [post]
func (h *UserFeedbackHandler) CreateUserFeedback(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	var req v1.CreateUserFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userFeedbackService.CreateUserFeedback(ctx, userCode, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
