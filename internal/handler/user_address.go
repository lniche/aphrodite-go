package handler

import (
	"aphrodite-go/internal/service"
	"github.com/gin-gonic/gin"
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

func (h *UserAddressHandler) GetUserAddress(ctx *gin.Context) {

}
