package repository

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserAddressRepository interface {
	GetUserAddress(ctx context.Context, userCode string, id uint64) (*model.UserAddress, error)
	GetUserAddresses(ctx context.Context, userCode string) (*[]model.UserAddress, error)
	CreateUserAddress(ctx context.Context, userAddress *model.UserAddress) error
	UpdateUserAddress(ctx context.Context, userAddress *model.UserAddress) error
	UpdateUserAddressDefault(ctx context.Context, userCode string, id uint64) error
	DeleteUserAddress(ctx context.Context, userCode string, id uint64) error
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

func (r userAddressRepository) GetUserAddress(ctx context.Context, userCode string, id uint64) (*model.UserAddress, error) {
	var userAddress model.UserAddress
	if err := r.DB(ctx).Where("user_code = ? and id = ?", userCode, id).First(&userAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &userAddress, nil
}

func (r userAddressRepository) GetUserAddresses(ctx context.Context, userCode string) (*[]model.UserAddress, error) {
	var userAddresses []model.UserAddress
	if err := r.DB(ctx).Where("user_code = ?", userCode).Find(&userAddresses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &userAddresses, nil
}

func (r userAddressRepository) CreateUserAddress(ctx context.Context, userAddress *model.UserAddress) error {
	if err := r.DB(ctx).Create(userAddress).Error; err != nil {
		return err
	}
	return nil
}

func (r userAddressRepository) UpdateUserAddress(ctx context.Context, userAddress *model.UserAddress) error {
	if err := r.DB(ctx).Save(userAddress).Error; err != nil {
		return err
	}
	return nil
}

func (r userAddressRepository) UpdateUserAddressDefault(ctx context.Context, userCode string, id uint64) error {
	if err := r.DB(ctx).Model(&model.UserAddress{}).Where("user_code = ? and id!=", userCode, id).Update("is_default", false).Error; err != nil {
		return err
	}
	return nil
}

func (r userAddressRepository) DeleteUserAddress(ctx context.Context, userCode string, id uint64) error {
	if err := r.DB(ctx).Where("user_code = ?", userCode).Delete(&model.UserAddress{}, id).Error; err != nil {
		return err
	}
	return nil
}
