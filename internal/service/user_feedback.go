package service

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
)

type UserFeedbackService interface {
	GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error)
	CreateUserFeedback(ctx context.Context, userCode string, req *v1.CreateUserFeedbackRequest) error
}

func NewUserFeedbackService(
	service *Service,
	userFeedbackRepository repository.UserFeedbackRepository,
	userRepository repository.UserRepository,
) UserFeedbackService {
	return &userFeedbackService{
		Service:                service,
		userFeedbackRepository: userFeedbackRepository,
		userRepository:         userRepository,
	}
}

type userFeedbackService struct {
	*Service
	userFeedbackRepository repository.UserFeedbackRepository
	userRepository         repository.UserRepository
}

func (s *userFeedbackService) GetUserFeedback(ctx context.Context, id int64) (*model.UserFeedback, error) {
	return s.userFeedbackRepository.GetUserFeedback(ctx, id)
}

func (s *userFeedbackService) CreateUserFeedback(ctx context.Context, userCode string, req *v1.CreateUserFeedbackRequest) error {
	_, err := s.userRepository.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	userFeedback := &model.UserFeedback{
		UserCode: userCode,
		Feedback: req.Feedback,
	}
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err = s.userFeedbackRepository.CreateUserFeedback(ctx, userFeedback); err != nil {
			return err
		}
		return nil
	})
	return nil
}
