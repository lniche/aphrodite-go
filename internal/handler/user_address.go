package handler

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserAddressHandler struct {
	*Handler
	userAddressService service.UserAddressService
}

func NewUserAddressHandler(
	handler *Handler,
	userAddressService service.UserAddressService,
) *UserAddressHandler {
	return &UserAddressHandler{
		Handler:            handler,
		userAddressService: userAddressService,
	}
}

// GetUserAddress godoc
// @Summary 获取地址详情
// @Schemes
// @Description
// @Tags 用户地址模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "地址ID"
// @Success 200 {object} v1.GetUserAddressResponse
// @Router /user/address/{id} [get]
func (h *UserAddressHandler) GetUserAddress(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	idStr := ctx.Param("id")
	if idStr == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userAddress, err := h.userAddressService.GetUserAddress(ctx, userCode, uint(id))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, userAddress)
}

// GetUserAddresses godoc
// @Summary 获取所有地址
// @Schemes
// @Description
// @Tags 用户地址模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetUserAddressesResponse
// @Router /user/address [get]
func (h *UserAddressHandler) GetUserAddresses(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)
	if userCode == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	userAddresses, err := h.userAddressService.GetUserAddresses(ctx, userCode)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, userAddresses)
}

// UpdateUserAddress godoc
// @Summary 修改用户地址
// @Schemes
// @Description
// @Tags 用户地址模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.UpdateUserAddressRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user/address [put]
func (h *UserAddressHandler) UpdateUserAddress(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	var req = new(v1.UpdateUserAddressRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userAddressService.UpdateUserAddress(ctx, userCode, req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// CreateUserAddress godoc
// @Summary 新增用户地址
// @Schemes
// @Description
// @Tags 用户地址模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.CreateUserAddressRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user/address [post]
func (h *UserAddressHandler) CreateUserAddress(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	var req = new(v1.CreateUserAddressRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userAddressService.CreateUserAddress(ctx, userCode, req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteUserAddress godoc
// @Summary 修改用户地址
// @Schemes
// @Description
// @Tags 用户地址模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "地址ID"
// @Success 200 {object} v1.Response
// @Router /user/address/{id} [delete]
func (h *UserAddressHandler) DeleteUserAddress(ctx *gin.Context) {
	userCode := GetUserCodeFromCtx(ctx)

	idStr := ctx.Param("id")
	if idStr == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.userAddressService.DeleteUserAddress(ctx, userCode, uint(id)); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
