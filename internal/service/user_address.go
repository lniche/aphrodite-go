package service

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
)

type UserAddressService interface {
	GetUserAddress(ctx context.Context, userCode string, id uint64) (*v1.GetUserAddressResponseData, error)
	GetUserAddresses(ctx context.Context, userCode string) (*[]v1.GetUserAddressResponseData, error)
	CreateUserAddress(ctx context.Context, userCode string, req *v1.CreateUserAddressRequest) error
	UpdateUserAddress(ctx context.Context, userCode string, req *v1.UpdateUserAddressRequest) error
	DeleteUserAddress(ctx context.Context, userCode string, id uint64) error
}

func NewUserAddressService(
	service *Service,
	userAddressRepository repository.UserAddressRepository,
	userRepo repository.UserRepository,
) UserAddressService {
	return &userAddressService{
		Service:               service,
		userAddressRepository: userAddressRepository,
		userRepo:              userRepo,
	}
}

type userAddressService struct {
	*Service
	userAddressRepository repository.UserAddressRepository
	userRepo              repository.UserRepository
}

func (s userAddressService) GetUserAddress(ctx context.Context, userCode string, id uint64) (*v1.GetUserAddressResponseData, error) {
	userAddress, err := s.userAddressRepository.GetUserAddress(ctx, userCode, id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserAddressResponseData{
		RecipientPhone:   userAddress.RecipientPhone,
		RecipientName:    userAddress.RecipientName,
		RecipientAddress: userAddress.RecipientAddress,
		Default:          userAddress.Default,
	}, nil
}

func (s userAddressService) GetUserAddresses(ctx context.Context, userCode string) (*[]v1.GetUserAddressResponseData, error) {
	userAddresses, err := s.userAddressRepository.GetUserAddresses(ctx, userCode)
	if err != nil {
		return nil, err
	}

	var responseData []v1.GetUserAddressResponseData
	for _, userAddress := range *userAddresses {
		responseData = append(responseData, v1.GetUserAddressResponseData{
			RecipientPhone:   userAddress.RecipientPhone,
			RecipientName:    userAddress.RecipientName,
			RecipientAddress: userAddress.RecipientAddress,
			Default:          userAddress.Default,
		})
	}

	return &responseData, nil
}

func (s userAddressService) CreateUserAddress(ctx context.Context, userCode string, req *v1.CreateUserAddressRequest) error {
	user, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	userAddress := &model.UserAddress{
		RecipientPhone:   req.RecipientPhone,
		RecipientAddress: req.RecipientAddress,
		RecipientName:    req.RecipientName,
		Default:          req.Default,
		UserCode:         user.UserCode,
	}
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err = s.userAddressRepository.CreateUserAddress(ctx, userAddress); err != nil {
			return err
		}
		if req.Default {
			if err = s.userAddressRepository.UpdateUserAddressDefault(ctx, userCode, userAddress.ID); err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (s userAddressService) UpdateUserAddress(ctx context.Context, userCode string, req *v1.UpdateUserAddressRequest) error {
	_, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	userAddress, err := s.userAddressRepository.GetUserAddress(ctx, userCode, req.ID)
	if err != nil {
		return err
	}
	userAddress.RecipientAddress = req.RecipientAddress
	userAddress.RecipientName = req.RecipientName
	userAddress.RecipientPhone = req.RecipientPhone
	userAddress.Default = req.Default
	if err = s.userAddressRepository.UpdateUserAddress(ctx, userAddress); err != nil {
		return err
	}
	if req.Default {
		if err = s.userAddressRepository.UpdateUserAddressDefault(ctx, userCode, userAddress.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s userAddressService) DeleteUserAddress(ctx context.Context, userCode string, id uint64) error {
	user, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}

	if err = s.userAddressRepository.DeleteUserAddress(ctx, user.UserCode, id); err != nil {
		return err
	}
	return nil
}
