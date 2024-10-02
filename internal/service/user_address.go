package service

import (
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
)

type UserAddressService interface {
	GetUserAddress(ctx context.Context, id int64) (*model.UserAddress, error)
}

func NewUserAddressService(
	service *Service,
	userAddressRepository repository.UserAddressRepository,
) UserAddressService {
	return &userAddressService{
		Service:               service,
		userAddressRepository: userAddressRepository,
	}
}

type userAddressService struct {
	*Service
	userAddressRepository repository.UserAddressRepository
}

func (s *userAddressService) GetUserAddress(ctx context.Context, id int64) (*model.UserAddress, error) {
	return s.userAddressRepository.GetUserAddress(ctx, id)
}
