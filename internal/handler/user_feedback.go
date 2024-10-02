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

// AddUserFeedback godoc
// @Summary 新增用户反馈
// @Schemes
// @Description
// @Tags 用户反馈模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.AddUserFeedbackRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user/feedback/add [post]
func (h *UserFeedbackHandler) AddUserFeedback(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	var req v1.AddUserFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userFeedbackService.AddUserFeedback(ctx, userCode, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
