package repository

import (
	"aphrodite-go/internal/model"
	"context"
)

type UserAddressRepository interface {
	GetUserAddress(ctx context.Context, id int64) (*model.UserAddress, error)
}

func NewUserAddressRepository(
	repository *Repository,
) UserAddressRepository {
	return &userAddressRepository{
		Repository: repository,
	}
}

type userAddressRepository struct {
	*Repository
}

func (r *userAddressRepository) GetUserAddress(ctx context.Context, id int64) (*model.UserAddress, error) {
	var userAddress model.UserAddress

	return &userAddress, nil
}
